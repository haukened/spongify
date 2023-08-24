import { Greet } from '$lib/wailsjs/go/main/App';
import type { PageLoad } from './$types';

export const load: PageLoad = async () => {
    const greet = await Greet('haukened');
    return {greeting: greet}
}