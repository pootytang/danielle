import { env } from '$env/dynamic/public';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = (async ({ fetch }) => {
  const res = await fetch(env.PUBLIC_HOST_URL + '/api/v1/public/shower_reveal')
  const items = await res.json()

  if (!res.ok) {
    console.log(`shower_reveal/+page.server.ts: Problem fetching ultrasound images. retrieved status: ${res.status}`)
    console.log(`shower_reveal/+page.server.ts: ${res.statusText}`)
    throw error(res.status, 'Ohh Boy, something happened!');
  }
  
  console.log('shower_reveal/+page.server.ts (PageServerLoad): returning the retrieved images')
  // console.table(items)

  return { images: items }
})