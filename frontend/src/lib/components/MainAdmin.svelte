<script lang="ts">
	import { USER_ADMIN } from '$lib/constants/UserTypeConstants';
	import Footer from '$lib/components/Footer.svelte';
	import Header from './Header.svelte';

	const { user } = $props();
</script>

<svelte:head>
	<title>Pagrindinis administratoriaus puslapis - Kavos Skrudykla</title>
</svelte:head>

<div class="flex min-h-screen flex-col">
	<Header userType={USER_ADMIN} />

	<article class="mt-top mb-bottom mx-page flex grow flex-col gap-4">
		<div class="rounded-md bg-white px-5 py-2.5">
			<h2 class="mb-5 text-4xl font-medium">Administratorius - {user.name}</h2>
			<p>Sveiki atvykę į "Kavos Skrudyklos" svetainę skirtą administratoriams.</p>
		</div>
		<div class="rounded-md bg-white px-5 py-2.5">
			<p class="mb-5">
				Peržiūrėti <a
					href="\users"
					class="font-medium text-yellow-500 underline hover:text-yellow-700 hover:no-underline"
					>vartotojus.</a
				>
			</p>
			<p>
				Peržiūrėti <a
					href="\orders"
					class="font-medium text-yellow-500 underline hover:text-yellow-700 hover:no-underline"
					>visus užsakymus.</a
				>
			</p>
		</div>
	</article>

	<Footer />
</div>
