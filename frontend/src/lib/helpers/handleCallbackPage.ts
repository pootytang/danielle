export default function getCallbackPage(queryParam: string, pathName: string) {
	let cbPage 
  if (!queryParam) {
    console.log(`getCallbackPage(): setting cbPage to path: ${pathName.split('/')[1]}`)
    cbPage = pathName.split('/')[1]
  } else {
    console.log(`getCallbackPage(): setting cbPage to search param: ${queryParam}`)
    cbPage = queryParam
  }

  if (!cbPage || cbPage === '/null' || cbPage === 'null') {
    console.log(`getCallbackPage(): callback is null, setting to home`)
    cbPage = 'home'
  }

  return cbPage
}