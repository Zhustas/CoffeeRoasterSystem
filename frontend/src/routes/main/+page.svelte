<script lang="ts">
	import { onDestroy, onMount } from 'svelte';

	let intersecting = $state(false);
	let header: HTMLElement;
	onMount(() => {
		if (typeof IntersectionObserver !== 'undefined') {
			const headerObserver = new IntersectionObserver(
				(entries) => {
					intersecting = entries[0].isIntersecting;
				},
				{
					root: null,
					threshold: 0,
					rootMargin: `${header.getBoundingClientRect().height}px`
				}
			);

			headerObserver.observe(header);
		}

		slides = slider.querySelectorAll('.slide');
		goToSlide(0);

		myInterval = setInterval(myIntervalFunction, 5000);
	});

	onDestroy(() => {
		clearInterval(myInterval);
	});

	let myInterval: number;
	function myIntervalFunction() {
		if (!hovering) {
			slideRight();
		}
	}

	$effect(() => {
		dots = dotsContainer.children;
	});

	// Slide efektas
	let slider: HTMLElement;
	let currentSlide = 0;
	let slides: NodeListOf<Element> = $state();
	let dots: HTMLCollection;
	let dotsContainer: HTMLElement;

	function goToSlide(slide: number) {
		slides.forEach((s, i) => {
			s.style.transform = `translateX(${100 * (i - slide)}%)`;
		});
	}

	function slideLeft() {
		currentSlide--;

		if (currentSlide === -1) {
			currentSlide = slides.length - 1;
		}

		goToSlide(currentSlide);
		activateDot(currentSlide);
	}

	function slideRight() {
		currentSlide++;

		if (currentSlide === slides.length) {
			currentSlide = 0;
		}

		goToSlide(currentSlide);
		activateDot(currentSlide);
	}

	function dotContainerOnClick(event: MouseEvent) {
		const target = event.target;
		if (target.dataset.slide) {
			handleDotClick(Number(target.dataset.slide) - 1);
		}
	}

	function handleDotClick(slide: number) {
		currentSlide = slide;
		goToSlide(slide);
		activateDot(slide);
	}

	function activateDot(slide: number) {
		for (let dot of dots) {
			dot.classList.remove('bg-gray-400', 'bg-white');
			if (Number(dot.dataset.slide) === slide + 1) {
				dot.classList.add('bg-gray-400');
			} else {
				dot.classList.add('bg-white');
			}
		}
	}

	let hovering = $state(false);
</script>

<header bind:this={header}>
	<nav
		class="flex h-16 border-b border-gray-400 bg-white px-64 {!intersecting
			? 'fixed top-0 z-50 w-full opacity-95'
			: ''}"
	>
		<a aria-label="Perkrauti šį puslapį" href="/main" class="self-center p-2">
			<img alt="Tekstas - Kavos Skrudykla" src="coffee_roaster_title.png" class="h-12" />
		</a>
		<div class="flex grow items-center justify-end font-medium">
			<a href="/main" class="mx-4 p-2">Užsakymai</a>
			<a href="/main" class="mx-4 p-2">Paskyra</a>
			<div class="h-full w-px bg-gray-400"></div>
			<a href="/main" class="mx-4 p-2">Krepšelis</a>
		</div>
	</nav>
</header>

<article class="mx-64 mt-5 flex flex-col">
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div
		onmouseenter={(hovering = true)}
		onmouseleave={(hovering = false)}
		id="recommendations"
		class="mx-10 flex flex-col items-center"
	>
		<div id="top" class="flex justify-center gap-10">
			<div class="z-10 flex size-12 self-center rounded-full bg-white text-center text-3xl">
				<button onclick={slideLeft} class="grow rounded-full">&larr;</button>
			</div>
			<div bind:this={slider} class="relative h-72 w-[52rem] overflow-hidden bg-white">
				<div class="slide absolute h-full w-full transition-transform duration-1000">
					<p>First</p>
				</div>
				<div class="slide absolute h-full w-full transition-transform duration-1000">
					<p>Second</p>
				</div>
				<div class="slide absolute h-full w-full transition-transform duration-1000">
					<p>Third</p>
				</div>
				<div class="slide absolute h-full w-full transition-transform duration-1000">
					<p>Fourth</p>
				</div>
			</div>
			<div class="z-10 flex size-12 self-center rounded-full bg-white text-center text-3xl">
				<button onclick={slideRight} class="grow rounded-full">&rarr;</button>
			</div>
		</div>
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<div
			id="dots"
			bind:this={dotsContainer}
			onclick={dotContainerOnClick}
			class="mt-3 flex grow gap-5"
		>
			{#each slides as _, i}
				<button
					aria-label="Slide dot {i + 1}"
					class="{currentSlide === i
						? 'bg-gray-400'
						: 'bg-white'} size-3 rounded-full border border-gray-300 shadow-lg shadow-black hover:bg-gray-400"
					data-slide={i + 1}
				></button>
			{/each}
		</div>
	</div>

	<hr class="mx-32 my-8 h-0.5 rounded border-0 bg-gray-300" />

	<div class="flex h-16 gap-5 rounded-md">
		<div id="filters" class="h-48 w-52 bg-white">
			<p class="p-5">Hi</p>
		</div>
		<div id="search-products" class="flex grow flex-col content-center gap-5">
			<div id="search" class="h-92 flex grow bg-white">
				<div class="ml-2 content-center">
					<input
						placeholder="Ieškoti"
						class="w-64 border-0 border-b-2 border-gray-400 focus:border-black focus:outline-none focus:ring-0"
					/>
				</div>
				<div class="grow"></div>
				<div class="mr-2 content-center">
					<select
						class="border-2 border-gray-400 hover:border-black focus:border-black focus:outline-none focus:ring-0"
					>
						<option value="lowest-price">Pigiausios viršuje</option>
						<option value="highest-price">Brangiausios viršuje</option>
					</select>
				</div>
			</div>
			<div id="products" class="h-96 bg-white">Products</div>
		</div>
	</div>
</article>
