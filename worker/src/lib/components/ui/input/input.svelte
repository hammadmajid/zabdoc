<script lang="ts">
	import type { HTMLInputAttributes, HTMLInputTypeAttribute } from "svelte/elements";
	import { cn, type WithElementRef } from "$lib/utils.js";

	type InputType = Exclude<HTMLInputTypeAttribute, "file">;

	type Props = WithElementRef<
		Omit<HTMLInputAttributes, "type"> &
			({ type: "file"; files?: FileList } | { type?: InputType; files?: undefined })
	>;

	let {
		ref = $bindable(null),
		value = $bindable(),
		type,
		files = $bindable(),
		class: className,
		...restProps
	}: Props = $props();
</script>

{#if type === "file"}
	<input
		bind:this={ref}
		data-slot="input"
		class={cn(
			"selection:bg-primary selection:text-primary-foreground neo-border-sm neo-shadow-sm placeholder:text-muted-foreground flex h-10 w-full min-w-0 bg-background px-3 pt-1.5 text-sm font-medium outline-none transition-all disabled:cursor-not-allowed disabled:opacity-50 md:text-sm focus:translate-x-[2px] focus:translate-y-[2px] focus:shadow-none",
			className
		)}
		type="file"
		bind:files
		bind:value
		{...restProps}
	/>
{:else}
	<input
		bind:this={ref}
		data-slot="input"
		class={cn(
			"bg-background selection:bg-primary selection:text-primary-foreground neo-border-sm neo-shadow-sm placeholder:text-muted-foreground flex h-10 w-full min-w-0 px-3 py-2 text-base outline-none transition-all disabled:cursor-not-allowed disabled:opacity-50 md:text-sm focus:translate-x-[2px] focus:translate-y-[2px] focus:shadow-none",
			className
		)}
		{type}
		bind:value
		{...restProps}
	/>
{/if}
