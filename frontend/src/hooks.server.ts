import type { Handle } from "@sveltejs/kit";
import { env } from '$env/dynamic/public';
import { getCallbackPage, privatePages } from "$helpers";

export const handle: Handle = async ({ event, resolve }) => {
  console.log("Handle Hook(): Called for path ", event.url.pathname)
  console.log("Handle Hook(): Called with param ", event.url.searchParams.get('page'))

  // HANDLE THE CALL BACK PAGE (add this to the helpers)
  const cbPage = getCallbackPage(event.url.searchParams.get('page') || '', event.url.pathname)
  // if (!event.url.searchParams.get('page')) {
  //   console.log(`Handle Hook(): setting cbPage to path: ${event.url.pathname.split('/')[1]}`)
  //   cbPage = event.url.pathname.split('/')[1]
  // } else {
  //   console.log(`Handle Hook(): setting cbPage to search param: ${event.url.searchParams.get('page')}`)
  //   cbPage = event.url.searchParams.get('page')
  // }

  // if (!cbPage || cbPage === '/null' || cbPage === 'null') {
  //   console.log(`Handle Hook(): callback is null, setting to home`)
  //   cbPage = 'home'
  // }
  console.log("Handle Hook(): Call Back page set to", cbPage)

  // HANDLE THE AUTH PAGES
  if (event.url.pathname.includes('/auth')) {
    console.log("Handle Hook(): auth page called. Going to page: ", event.url.pathname)
    return await resolve(event)
  }

  // IF USER IS ALREADY LOGGED IN CONTINUE TO PAGE
  if (event.locals.user) {
    console.log(`Handle Hook(): User id ${event.locals.user.id} is already logged in. Load Page as normal`)
    return await resolve(event);
  }
  // WE DO NOT HAVE A LOCAL USER

  // IF WE HAVE AN ACCESS TOKEN, GET THE USER THAT IT BELONGS TO
  const at = event.cookies.get('access_token')
  if (at) {
    console.log("Handle Hook(): Access Token with no locals. calling the API's user endpoint")
    const res = await event.fetch(env.PUBLIC_HOST_URL + '/api/v1/auth/user', {
      credentials: 'include',
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ access_token: at })
    })

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
    
    // Result was ok set the locals
    const response = await res.json();
    console.log('Handle Hook(): retrieved user. Setting locals')
    event.locals.user = response.user
    console.log(`Handle Hook(): resolving to url: ${event.url.pathname} with params ${event.url.searchParams}`)
    return await resolve(event)
  }

  // NO LOCAL USER AND NO ACCESS TOKEN

  // LOG USER IN IF PAGE IS PRIVATE
  
  if (privatePages.includes(event.url.pathname)) {
    console.log('Handle Hook(): private page called. authorization required')
    return new Response('Redirect', {status: 303, headers: { Location: `/auth/login?page=${cbPage}`}})
  }

  // NO LOCAL USER, NO ACCESS TOKEN, PUBLIC PAGE REQUESTED. ACCESS IS ALLOWED (MAYBE ALWAYS REQUIRE A LOGON)
  console.log('Handle Hook(): Public page called without logged in user.  ALL GOOD')
  return await resolve(event)
}