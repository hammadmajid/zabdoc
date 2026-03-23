<script lang="ts">
	import { wizardStore } from "$lib/stores/wizard-store.svelte";
	import User from "@lucide/svelte/icons/user";
	import Users from "@lucide/svelte/icons/users";
	import FileQuestion from "@lucide/svelte/icons/file-question";
</script>

<div class="flex flex-col items-center justify-center px-4 py-8 text-center">
	<div class="neo-border neo-shadow-lg mb-8 rotate-1 bg-accent px-8 py-4">
		<h1 class="text-3xl font-black tracking-tight uppercase md:text-4xl">
			Working alone or in a group?
		</h1>
	</div>

	<p class="mb-8 max-w-md text-lg font-medium">
		This helps us know whose information to collect.
	</p>

	<div class="flex flex-col gap-6 sm:flex-row">
		<button
			type="button"
			onclick={() => {
				wizardStore.setTeamType("individual");
				wizardStore.nextStep();
			}}
			class="neo-border neo-shadow flex cursor-pointer flex-col items-center gap-4 bg-card px-10 py-8 text-xl font-black tracking-wide uppercase transition-all hover:translate-x-1 hover:translate-y-1 hover:bg-secondary hover:shadow-none"
		>
			<div class="neo-border-sm bg-primary p-4">
				<User class="size-10" />
			</div>
			<span>Just Me</span>
		</button>
		<button
			type="button"
			onclick={() => {
				wizardStore.setTeamType("group");
				wizardStore.nextStep();
			}}
			class="neo-border neo-shadow flex cursor-pointer flex-col items-center gap-4 bg-card px-10 py-8 text-xl font-black tracking-wide uppercase transition-all hover:translate-x-1 hover:translate-y-1 hover:bg-accent hover:shadow-none"
		>
			<div class="neo-border-sm bg-secondary p-4">
				<Users class="size-10" />
			</div>
			<span>Group</span>
		</button>
	</div>

	<button
		type="button"
		onclick={() => {
			wizardStore.setTeamType("blank");
			wizardStore.nextStep();
		}}
		class="neo-border neo-shadow mt-6 flex cursor-pointer items-center gap-4 bg-muted px-6 py-4 transition-all hover:translate-x-0.5 hover:translate-y-0.5 hover:bg-card hover:shadow-none"
	>
		<div class="neo-border-sm bg-card p-2">
			<FileQuestion class="size-6" />
		</div>
		<div class="text-left">
			<span class="font-bold uppercase">Leave Blank</span>
			<p class="text-sm font-normal text-muted-foreground normal-case">
				Fill in name & reg number by hand after printing
			</p>
		</div>
	</button>
</div>
