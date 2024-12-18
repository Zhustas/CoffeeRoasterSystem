<script lang="ts">
	import { USER_ROASTER } from '$lib/constants/UserTypeConstants';
	import Footer from '../../components/Footer.svelte';
	import Header from './Header.svelte';

	const { user } = $props();
</script>

<div class="flex min-h-screen flex-col">
	<Header userType={USER_ROASTER} />

	<article class="mt-top mb-bottom mx-page grow">
		<div class="rounded-md bg-white px-5 py-2.5">
			<h2 class="mb-5 text-4xl font-medium">Darbuotojas - {user.name}</h2>
			<p>Sveiki atvykę į "Kavos Skrudyklos" svetainę skirtą darbuotojams.</p>
		</div>
	</article>

	<Footer />
</div>
