<script lang="ts">
	import Header from '$lib/components/Header.svelte';
	import { USER_ADMIN, USER_CUSTOMER, USER_ROASTER } from '$lib/constans';
	import Footer from '../../components/Footer.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
	let user = data.user;
</script>

<svelte:head>
	<title>DUK - Kavos Skrudykla</title>
</svelte:head>

<div class="flex min-h-screen flex-col">
	{#if user.role === USER_CUSTOMER}
		<Header userType={USER_CUSTOMER} />
	{:else if user.role === USER_ROASTER}
		<Header userType={USER_ROASTER} />
	{:else if user.role === USER_ADMIN}
		<Header userType={USER_ADMIN} />
	{/if}

	<article class="mt-top mb-bottom mx-page grow">
		<h2 class="text-2xl font-medium">Sveiki atvykę į DUK puslapį!</h2>
	</article>

	<Footer />
</div>
