<script lang="ts">
	import Header from '$lib/components/Header.svelte';
	import { USER_NONE } from '$lib/constants/UserTypeConstants';
	import Footer from '$lib/components/Footer.svelte';
</script>

<div class="flex min-h-screen flex-col">
	<Header userType={USER_NONE} />

	<article class="mx-page mt-top mb-bottom grow">
		<div class="ml-2">
			<h2 class="mb-5 text-4xl font-medium text-red-500">Puslapis nerastas. (404)</h2>
			<p class="text-lg">
				Eikite į <a href="/main" class="font-medium underline">pagrindinį puslapį</a> arba
				<a href="/" class="font-medium underline">prisijunkite</a>. Jeigu neturite paskyros,
				<a href="/" class="font-medium underline">prisiregistruokite</a>.
			</p>
		</div>
	</article>

	<Footer />
</div>
