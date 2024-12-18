<script lang="ts">
	import { onMount } from 'svelte';

	let {
		id,
		name,
		price,
		roastType
	}: { id: number; name: string; price: string; roastType: string } = $props();

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
		{#if roastType === 'light'}
			<img alt="Light roast coffee bean" src="light_bean.png" class="w-full cursor-pointer" />
		{:else if roastType === 'medium'}
			<img alt="Medium roast coffee bean" src="medium_bean.png" class="w-full cursor-pointer" />
		{:else if roastType === 'dark'}
			<img alt="Dark roast coffee bean" src="dark_bean.png" class="w-full cursor-pointer" />
		{/if}
	</a>
	<p id="title" class="mt-2 w-fit cursor-text">{name}</p>
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
