import { env } from '$env/dynamic/public';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = (async ({ fetch }) => {
  const res = await fetch(env.PUBLIC_HOST_URL + '/api/v1/public/growingmommy')
  const items = await res.json()

  if (!res.ok) {
    console.log(`growing_mommy/+page.server.ts: Problem fetching ultrasound images. retrieved status: ${res.status}`)
    console.log(`growing_mommy/+page.server.ts: ${res.statusText}`)
    throw error(res.status, 'Ohh Boy, something happened!');
  }
  
  console.log('growing_mommy/+page.server.ts (PageServerLoad): returning the retrieved images')
  // console.table(items)

  return { images: items }
})