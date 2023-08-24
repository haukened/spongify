import { goto } from '$app/navigation'
import { isValidURL } from '$lib/utils'
import { EventsOn, WindowReloadApp } from '$lib/wailsjs/runtime/runtime' 

// this force reloads the app on error
export async function handleError() {
    WindowReloadApp()
}

// Subscribe the client frontend to navigation events
EventsOn("navigate", doNavigation)

// handler for navigation event
function doNavigation(data: any) {
    if (data && isValidURL(data)) {
        goto(data)
    } else {
        console.log("invalid URI for navigation: ", data);
    }
}