<script lang="ts">
	import { USER_ROASTER } from '$lib/constants/UserTypeConstants';
	import Footer from '$lib/components/Footer.svelte';
	import Header from './Header.svelte';

	const { user } = $props();
</script>

<div class="flex min-h-screen flex-col">
	<Header userType={USER_ROASTER} />

	<article class="mx-page mb-bottom mt-top flex grow flex-col gap-4">
		<div class="rounded-md bg-white px-5 py-2.5">
			<h2 class="mb-5 text-4xl font-medium">Darbuotojas - {user.name}</h2>
			<p>Sveiki atvykę į "Kavos Skrudyklos" svetainę skirtą darbuotojams.</p>
		</div>
		<div class="rounded-md bg-white px-5 py-2.5">
			<p class="mb-5">
				Peržiūrėti <a
					href="\inventory"
					class="font-medium text-yellow-500 underline hover:text-yellow-700 hover:no-underline"
					>kavos inventorių.</a
				>
			</p>
			<p class="mb-5">
				Peržiūrėti <a
					href="\orders"
					class="font-medium text-yellow-500 underline hover:text-yellow-700 hover:no-underline"
					>visus užsakymus.</a
				>
			</p>
			<p>
				Peržiūrėti <a
					href="\neworders"
					class="font-medium text-yellow-500 underline hover:text-yellow-700 hover:no-underline"
					>naujus užsakymus.</a
				>
			</p>
		</div>
	</article>

	<Footer />
</div>
