<script lang="ts">
	import { onMount } from 'svelte';

	let { id, name, price }: { id: number; name: string; price: string } = $props();

	let euroFormatter: Intl.NumberFormat = $state();

	onMount(() => {
		euroFormatter = new Intl.NumberFormat(navigator.language, {
			style: 'currency',
			currency: 'EUR'
		});
	});
</script>

<div class="flex w-52 flex-col rounded-md bg-white p-2 hover:shadow-md hover:shadow-gray-400">
	<a href="/coffee/{id}">
		<img alt="Coffee bean" src="coffee_bean.jpg" class="cursor-pointer" />
	</a>
	<p id="title" class="mt-1 w-fit cursor-text text-sm">{name}</p>
	<div class="flex justify-end">
		<p id="price" class="mt-4 w-fit cursor-text break-words text-end text-lg font-medium">
			{#if euroFormatter}
				{euroFormatter.format(Number(price))}
			{:else}
				Kaina kraunama...
			{/if}
		</p>
	</div>
</div>
