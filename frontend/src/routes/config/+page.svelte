<script lang="ts">
	import { Hotkey } from "$lib/components/hotkeys";

    let keys: string[] = ["ctrl", "shift", "m"]
    let keysListening: boolean = false;

    const listenToKeys = () => {

    }
</script>

<div class="flex flex-col items-left space-y-2">
    <h1 class="text-2xl font-bold underline">Preferences:</h1>
    <label class="flex flex-row space-x-4 items-center">
        <input type="checkbox" class="toggle toggle-sm toggle-success"/>
        <span>Start on login</span>
    </label>
    <span>HotKeys</span>
    <Hotkey/>
</div>