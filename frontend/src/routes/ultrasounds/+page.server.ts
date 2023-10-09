import { env } from '$env/dynamic/public';
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = (async ({ fetch }) => {
  console.log("Fetching: " + env.PUBLIC_HOST_URL + '/api/v1/public/ultrasounds')
  const res = await fetch(env.PUBLIC_HOST_URL + '/api/v1/public/ultrasounds')
  const items = await res.json()

  if (!res.ok) {
    console.log(`Ultrasounds/+page.server.ts (PageServerLoad): Problem fetching ultrasound images. retrieved status: ${res.status}`)
    console.log(`Ultrasounds/+page.server.ts (PageServerLoad): ${res.statusText}`)
    throw error(res.status, 'Ohh Boy, something happened!');
  }
  
  console.log('Ultrasounds/+page.server.ts (PageServerLoad): returning the retrieved images')
  return { images: items }
})