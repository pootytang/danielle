import { env } from '$env/dynamic/public';
import { fail, redirect } from '@sveltejs/kit';
import type { Actions } from "./$types";
import type { PageServerLoad } from './$types';
import { redirect2CallbackPage } from '$helpers';

export const load: PageServerLoad = (async ({ url, locals }) => {
  console.log(`signup/+page.server.ts (PageServerLoad): URL Params: ${url.searchParams.get('page')}`)

  let cbPage = url.searchParams.get('page')
  if (!cbPage || cbPage === 'null') {
    console.log(`signup/+page.server.ts (PageServerLoad): setting cbPage to path: /`)
    cbPage = '/'
  } 
  console.log("signup/+page.server.ts (PageServerLoad): Call Back page set to", cbPage)

  console.log('signup/+page.server.ts (PageServerLoad): Checking if user is already logged in')
  if (locals.user) {
    if (cbPage === '/') {
      throw redirect(302, cbPage)
    } else {
      throw redirect(302, `/${cbPage}`)
    }
  }

  console.log('signup/+page.server.ts (PageServerLoad): User is NOT logged in')
  return { data: cbPage }
})

export const actions: Actions = {
  default: async ({request,url, fetch, cookies}) => {
    console.log(`signup/+page.server.ts (Actions): URL Params: ${url.searchParams.get('page')}`)
    
    let page = url.searchParams.get('page')
    if (!page || page === 'null') {
      console.log('signup/+page.server.ts (Actions): setting callback page to "home"')
      page = 'home'
    }
    const formData = Object.fromEntries(await request.formData());

    console.log("signup/+page.server.ts (Actions): Checking that name, email, and password was populated")
    if (!formData.email || !formData.password || !formData.name) {
			return fail(400, {
				error: 'Missing name, email or password'
			});
		}
    
    console.log("signup/+page.server.ts (Actions): Posting to API")
    const res = await fetch(env.PUBLIC_HOST_URL + '/api/v1/auth/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        name: formData.name,
        email: formData.email,
        password: formData.pass,
        passwordConfirm: formData.cconfirm_password
      })
    })
    
    console.log("signup/+page.server.ts (Actions): Got result of: ", res.status)
    if (res.status != 201) {
      console.error("signup/+page.server.ts (Actions): Unfavorable result")
      return fail(res.status, {
        error: "server responded with an error"
      })
    }

    console.log("signup/+page.server.ts (Actions): Signed up successfully. Grab access token")
    const jsonData = await res.json()

    // User object shape: (ID, Name, Email, Access_Token, Refresh_Token, Logged_In, Role, CreatedAt, UpdatedAt)
    const access_token = jsonData.user.access_token
    const refresh_token = jsonData.user.refresh_token

    cookies.set('access_token', access_token, {
      httpOnly: true,
      path: '/',
      secure: false,
      maxAge: 60 * 60 * 24 // 1 day
    })

    cookies.set('refresh_token', refresh_token, {
      httpOnly: true,
      path: '/',
      secure: false,
      maxAge: 60 * 60 * 24 * 5 // 5 days
    })

    console.log(`singup/+page.server.ts (Actions): Access and Refresh tokens set in cookies. Redirecting to ${page}`)
    throw redirect2CallbackPage(302, page)
    // if (page === 'home') {
    //   throw redirect(302, '/')
    // } else {
    //   throw redirect(302, `/${page}`)
    // }
  }
};