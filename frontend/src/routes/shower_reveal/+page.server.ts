import { env } from '$env/dynamic/public';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = (async ({ fetch }) => {
  const res = await fetch(env.PUBLIC_HOST_URL + '/api/v1/public/shower_reveal')
  const items = await res.json()

  console.log("shower_reveal/+page.server.ts:")
  console.table(items)

  return { images: items }
})