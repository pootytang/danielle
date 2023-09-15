import { env } from '$env/dynamic/public';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = (async ({ fetch }) => {
  console.log("Fetching: " + env.PUBLIC_HOST_URL + '/api/v1/public/ultrasounds')
  const res = await fetch(env.PUBLIC_HOST_URL + '/api/v1/public/ultrasounds')
  const items = await res.json()

  console.log("Ultrasounds/+page.server.ts:")
  console.table(items)

  return { images: items }
})