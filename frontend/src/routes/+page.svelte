<script>
	import { goto } from "$app/navigation";
</script>
<p>Oh hai! How did you get here?</p>
<button class="btn btn-primary" on:click={() => goto("/about")}>Go Here Instead</button>