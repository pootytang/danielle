import type { Handle } from "@sveltejs/kit";
import { env } from '$env/dynamic/public';

export const handle: Handle = async ({ event, resolve }) => {
  console.log("Handle Hook(): Called for path ", event.url.pathname)
  console.log("Handle Hook(): Called with param ", event.url.searchParams.get('page'))

  // HANDLE THE CALL BACK PAGE
  let cbPage 
  if (!event.url.searchParams.get('page')) {
    console.log(`Handle Hook(): setting cbPage to path: ${event.url.pathname.split('/')[1]}`)
    cbPage = event.url.pathname.split('/')[1]
  } else {
    console.log(`Handle Hook(): setting cbPage to search param: ${event.url.searchParams.get('page')}`)
    cbPage = event.url.searchParams.get('page')
  }

  if (!cbPage || cbPage === '/null' || cbPage === 'null') {
    console.log(`Handle Hook(): callback is null, setting to /`)
    cbPage = '/'
  }
  console.log("Handle Hook(): Call Back page set to", cbPage)

  // HANDLE REFRESH PAGE
  if (event.url.pathname.includes('/refresh')) {
    console.log('Handle Hook(): going to refresh page')
    return await resolve(event)
  }

  // const private_pages = ['/delivery', '/dob']
  // if (!private_pages.includes(event.url.pathname)) {
  //   console.log('Handle Hook(): public page called. continue to page')
  //   return await resolve(event)
  // }
  // console.log(`Handle Hook(): Page ${event.url.pathname} is a private page`)
  
  if (event.locals.user) {
    console.log(`Handle Hook(): User id ${event.locals.user.id} is already logged in. Load Page as normal`)
    return await resolve(event);
	}

  const at = event.cookies.get('access_token')
  if (!at) {
    console.log(`Handle Hook(): There's no access token found. Load Page as normal`)
    return await resolve(event);
  }

  // We have an access token but no local user. Request the user endpoint to get the user belonging to the token and store it in locals.user
  if (at) {
    console.log("Handle Hook(): Access Token with no locals. calling the API's user endpoint")
    const res = await event.fetch(env.PUBLIC_HOST_URL + '/api/v1/auth/user', {
      credentials: 'include',
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        access_token: at
      })
    });

    if (!res.ok) {
      switch (res.status) {
        case 400: // sent bad data
          console.log('Handle Hook(): received a 400 status code, bad access token sent to server. Redirecting to login page')
          return new Response('Redirect', {status: 303, headers: { Location: `/auth/login?page=${cbPage}`}})
          break;
        case 401: // unathorized, expired token
          console.log('Handle Hook(): received a 401 status code, expired token. Redirecting to refresh page')
          return new Response('Redirect', {status: 303, headers: { Location: `/auth/refresh?page=${cbPage}`}})
          break;
        case 403: // forbidden, can't find the token subject
          console.log('Handle Hook(): received a 403 status code, subject not found in token. Redirecting to login page')
          return new Response('Redirect', {status: 303, headers: { Location: `/auth/login?page=${cbPage}`}})
          break;
        default: // may want to throw an error and show a sveltekit error page here.
          console.log(`Handle Hook(): received unhandled status code ${res.status}`)
          return await resolve(event)
      }
    }

    const response = await res.json();
    console.log('Handle Hook(): retrieved user. Setting locals')
    event.locals.user = response.user
  }

  const response = await resolve(event)
  console.log("Handle Hook(): Returning Response")
  return response
}