<script lang="ts">
	import { Switch as SwitchPrimitive } from "bits-ui";
	import { cn, type WithoutChildrenOrChild } from "$lib/utils.js";

	let {
		ref = $bindable(null),
		class: className,
		checked = $bindable(false),
		...restProps
	}: WithoutChildrenOrChild<SwitchPrimitive.RootProps> = $props();
</script>

<SwitchPrimitive.Root
	bind:ref
	bind:checked
	data-slot="switch"
	class={cn(
		"data-[state=checked]:bg-primary data-[state=unchecked]:bg-muted neo-border-sm peer inline-flex h-6 w-11 shrink-0 items-center outline-none transition-all disabled:cursor-not-allowed disabled:opacity-50",
		className
	)}
	{...restProps}
>
	<SwitchPrimitive.Thumb
		data-slot="switch-thumb"
		class={cn(
			"bg-foreground pointer-events-none block size-4 ring-0 transition-transform data-[state=checked]:translate-x-[calc(100%+4px)] data-[state=unchecked]:translate-x-1"
		)}
	/>
</SwitchPrimitive.Root>
