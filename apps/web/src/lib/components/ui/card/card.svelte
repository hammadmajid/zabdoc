<script lang="ts">
	import type { HTMLAttributes } from "svelte/elements";
	import { cn, type WithElementRef } from "$lib/utils.js";

	let {
		ref = $bindable(null),
		class: className,
		children,
		...restProps
	}: WithElementRef<HTMLAttributes<HTMLDivElement>> = $props();
</script>

<div
	bind:this={ref}
	data-slot="card"
	class={cn(
		"neo-border neo-shadow flex flex-col gap-6 bg-card py-6 text-card-foreground",
		className
	)}
	{...restProps}
>
	{@render children?.()}
</div>
