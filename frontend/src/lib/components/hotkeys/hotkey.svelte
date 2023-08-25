<script lang="ts">
    export let key: string = 'M';
    export let modifiers: string[] = ['Ctrl', 'Shift'];
    let isFocused = false;

    // we need to do this becuase svelte's reactivity is based on assignment
    // so we can't just use array.push(mod) because svelte doens't see the assignment
    function addModifier(mod: string) {
        if (!modifiers.includes(mod)) {
            modifiers = [...modifiers, mod]
        }
    }

    function removeModifier(mod: string) {
        let filtered = modifiers.filter(function(value){
            return value !== mod;
        })
        modifiers = filtered;
    }

    const handleKeyDown = (event: KeyboardEvent) => {
        // return if this is just a repeated key
        if (event.repeat) {
            return
        }
        // check if this is just a mod key, if so wait for the final keycombo
        if (
            (!event.altKey && !event.ctrlKey && !event.metaKey && !event.shiftKey) ||
            event.key === 'Meta' ||
            event.key === 'Shift' ||
            event.key === 'Control' ||
            event.key === 'Alt'
        ) {
            return
        }
        key = event.key.toUpperCase();
        if (event.metaKey) {
            addModifier("Meta")
        } else {
            removeModifier('Meta')
        }
        if (event.ctrlKey) {
            addModifier("Ctrl")
        } else {
            removeModifier('Ctrl')
        }
        if (event.shiftKey) {
            addModifier('Shift')
        } else {
            removeModifier('Shift')
        }
        if (event.altKey) {
            addModifier('Alt')
        } else {
            removeModifier('Alt')
        }
        isFocused = false;
    }
</script>

<div class="join">
    <div 
        role="textbox" 
        tabindex="0" 
        on:focusin={() => isFocused = true}
        on:focusout={() => isFocused = false} 
        on:keydown|preventDefault={handleKeyDown} 
        class="input input-bordered no-ring cursor-pointer items-center justify-center w-52 flex flex-row space-x-1 p-2 {isFocused? 'input-primary' : ''}"
    >
    {#if isFocused}
        <div>Listening for keys...</div>
    {:else}
        {#each modifiers as mod, idx }
            <div id="key-{idx}" class="kbd px-2">{mod}</div>
            <span>+</span>
        {/each}
        {#if key !== ''}
        <div class="kbd px-2">{key}</div>
        {/if}
    {/if}
    </div>
</div>