<script lang="ts">
	import Header from '$lib/components/Header.svelte';
	import { USER_CUSTOMER } from '$lib/constans';
	import Footer from '../../components/Footer.svelte';
</script>

<svelte:head>
	<title>Prenumerata - Kavos Skrudykla</title>
</svelte:head>

<div class="flex min-h-screen flex-col">
	<Header userType={USER_CUSTOMER} />

	<article class="mx-page mt-top mb-bottom grow">
		<h2 class="text-2xl font-medium">Sveiki atvykę į prenumeratą!</h2>
	</article>

	<Footer />
</div>
