<script lang="ts">
	import Header from '$lib/components/Header.svelte';
	import { USER_CUSTOMER } from '$lib/constants/UserTypeConstants';
	import Footer from '$lib/components/Footer.svelte';
</script>

<svelte:head>
	<title>Prenumerata - Kavos Skrudykla</title>
</svelte:head>

<div class="flex min-h-screen flex-col">
	<Header userType={USER_CUSTOMER} />

	<article class="mx-page mb-bottom mt-top grow">
		<h2 class="text-2xl font-medium">Sveiki atvykę į prenumeratą!</h2>
	</article>

	<Footer />
</div>
