<script lang="ts">
	import * as AlertMessageConstants from '../constants/AlertMessageConstants';

	let { status, message }: { status: string; message: string } = $props();

	const COLORS = {
		[AlertMessageConstants.STATUS_SUCCESS]: 'bg-green-500',
		[AlertMessageConstants.STATUS_INFORMATION]: 'bg-yellow-500',
		[AlertMessageConstants.STATUS_FAILURE]: 'bg-red-500'
	};
</script>

<div
	class="animate-slide-message absolute {COLORS[
		status
	]} right-5 top-16 px-8 py-4 shadow-md shadow-black"
>
	<p class="font-medium text-white">{message}</p>
</div>
