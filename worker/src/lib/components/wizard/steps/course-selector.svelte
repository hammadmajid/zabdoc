<script lang="ts">
	import { RadioGroup, RadioGroupItem } from "$lib/components/ui/radio-group/index.js";
	import { cn } from "$lib/utils.js";
	import Calculator from "@lucide/svelte/icons/calculator";
	import GitBranch from "@lucide/svelte/icons/git-branch";
	import HardDrive from "@lucide/svelte/icons/hard-drive";
	import Wrench from "@lucide/svelte/icons/wrench";
	import Code from "@lucide/svelte/icons/code";
	import Beaker from "@lucide/svelte/icons/beaker";

	interface Course {
		value: string;
		label: string;
		icon?: string;
	}

	const iconMap: Record<string, any> = {
		calculator: Calculator,
		"git-branch": GitBranch,
		"hard-drive": HardDrive,
		wrench: Wrench,
		code: Code,
		beaker: Beaker,
	};

	let { value = $bindable(""), courses = [], disabled = false }: { value: string; courses: Course[]; disabled: boolean } = $props();

	function getIconComponent(iconName?: string) {
		if (!iconName || !iconMap[iconName]) {
			return null;
		}
		return iconMap[iconName];
	}
</script>

<RadioGroup bind:value {disabled} class="flex flex-col gap-3 w-full">
	{#each courses as course (course.value)}
		{@const IconComponent = getIconComponent(course.icon)}
		<label
			class={cn(
				"neo-border neo-shadow px-6 py-4 flex items-center gap-3 font-bold uppercase cursor-pointer transition-all",
				disabled && "opacity-50 cursor-not-allowed",
				value === course.value 
					? "bg-accent text-accent-foreground" 
					: "bg-card hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none"
			)}
		>
			<RadioGroupItem value={course.value} class="sr-only" {disabled} />
			{#if IconComponent}
				<IconComponent class="size-5 flex-shrink-0" />
			{/if}
			{course.label}
		</label>
	{/each}
</RadioGroup>
