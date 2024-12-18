<script lang="ts">
	import { USER_ADMIN } from '$lib/constants/UserTypeConstants';
	import Footer from '../../components/Footer.svelte';
	import Header from './Header.svelte';

	const { user } = $props();
</script>

<svelte:head>
	<title>Pagrindinis administratoriaus puslapis - Kavos Skrudykla</title>
</svelte:head>

<div class="flex min-h-screen flex-col">
	<Header userType={USER_ADMIN} />

	<article class="mt-top mb-bottom mx-page grow">
		<div class="rounded-md bg-white px-5 py-2.5">
			<h2 class="mb-5 text-4xl font-medium">Administratorius - {user.name}</h2>
			<p>Sveiki atvykę į "Kavos Skrudyklos" svetainę skirtą administratoriams.</p>
		</div>
	</article>

	<Footer />
</div>
