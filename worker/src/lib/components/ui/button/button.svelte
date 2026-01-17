<script lang="ts" module>
	import { cn, type WithElementRef } from "$lib/utils.js";
	import type { HTMLAnchorAttributes, HTMLButtonAttributes } from "svelte/elements";
	import { type VariantProps, tv } from "tailwind-variants";

	export const buttonVariants = tv({
		base: "neo-border-sm neo-shadow-sm inline-flex shrink-0 items-center justify-center gap-2 whitespace-nowrap text-sm font-bold uppercase tracking-wide outline-none transition-all hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none focus-visible:translate-x-[2px] focus-visible:translate-y-[2px] focus-visible:shadow-none disabled:pointer-events-none disabled:opacity-50 aria-disabled:pointer-events-none aria-disabled:opacity-50 [&_svg:not([class*='size-'])]:size-4 [&_svg]:pointer-events-none [&_svg]:shrink-0",
		variants: {
			variant: {
				default: "bg-primary text-primary-foreground",
				destructive: "bg-destructive text-white",
				outline: "bg-background text-foreground",
				secondary: "bg-secondary text-secondary-foreground",
				ghost: "border-transparent shadow-none hover:bg-muted hover:translate-x-0 hover:translate-y-0 hover:shadow-none",
				link: "text-primary underline-offset-4 hover:underline border-transparent shadow-none hover:translate-x-0 hover:translate-y-0",
			},
			size: {
				default: "h-10 px-5 py-2 has-[>svg]:px-4",
				sm: "h-9 gap-1.5 px-4 has-[>svg]:px-3",
				lg: "h-12 px-8 text-base has-[>svg]:px-6",
				icon: "size-10",
			},
		},
		defaultVariants: {
			variant: "default",
			size: "default",
		},
	});

	export type ButtonVariant = VariantProps<typeof buttonVariants>["variant"];
	export type ButtonSize = VariantProps<typeof buttonVariants>["size"];

	export type ButtonProps = WithElementRef<HTMLButtonAttributes> &
		WithElementRef<HTMLAnchorAttributes> & {
			variant?: ButtonVariant;
			size?: ButtonSize;
		};
</script>

<script lang="ts">
	let {
		class: className,
		variant = "default",
		size = "default",
		ref = $bindable(null),
		href = undefined,
		type = "button",
		disabled,
		children,
		...restProps
	}: ButtonProps = $props();
</script>

{#if href}
	<a
		bind:this={ref}
		data-slot="button"
		class={cn(buttonVariants({ variant, size }), className)}
		href={disabled ? undefined : href}
		aria-disabled={disabled}
		role={disabled ? "link" : undefined}
		tabindex={disabled ? -1 : undefined}
		{...restProps}
	>
		{@render children?.()}
	</a>
{:else}
	<button
		bind:this={ref}
		data-slot="button"
		class={cn(buttonVariants({ variant, size }), className)}
		{type}
		{disabled}
		{...restProps}
	>
		{@render children?.()}
	</button>
{/if}
