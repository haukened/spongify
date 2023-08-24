<script lang="ts">
	import { BrowserOpenURL, WindowHide } from "$lib/wailsjs/runtime/runtime";

    const spongeCase = (text: string): string => {
        return text.split('').map(l => Math.random() < 0.5 && l.toUpperCase() || l.toLowerCase()).join('')
    }
    const description = spongeCase("Add sarcasm to every conversation.")

    const handleClick = (url: string) => {
        try {
            const parsed = new URL(url)
            BrowserOpenURL(parsed.toString())
            WindowHide()
        } 
        catch(e) {
            console.log(e)
        }
    }
</script>

<div class="flex flex-col m-4 max-w-fit items-center">
    <div class="flex flex-row items-center justify-between w-full">
        <h1 class="text-2xl underline font-bold">Spongify!</h1>
        <img class="max-h-[50px] ml-10" src="/favicon.png" alt="spongebob mocking you"/>
    </div>
    <p class="text-sm font-bold">{description}</p>
    <div class="flex flex-col items-center">
        <p class="mt-4 text-lg">Made by David Haukeness</p>
        <div class="flex flex-row mt-4 space-x-2">
            <button class="btn btn-sm btn-primary" on:click={() => handleClick("https://github.com/haukened/spongify")}><i class="fa-brands fa-github"/>github</button>
            <div class="divider"/>
            <button class="btn btn-sm btn-warning" on:click={() => handleClick("https://www.buymeacoffee.com/haukened")}>Buy me a <i class="fa fa-beer"/></button>
        </div>
    </div>
</div>