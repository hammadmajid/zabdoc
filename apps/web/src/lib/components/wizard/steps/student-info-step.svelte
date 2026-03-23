<script lang="ts">
	import { wizardStore } from "$lib/stores/wizard-store.svelte";
	import { formStore, STUDENT_LIMIT } from "$lib/stores/form-store.svelte";
	import Input from "$lib/components/ui/input/input.svelte";
	import Label from "$lib/components/ui/label/label.svelte";
	import Button from "$lib/components/ui/button/button.svelte";
	import Plus from "@lucide/svelte/icons/plus";
	import X from "@lucide/svelte/icons/x";
	import User from "@lucide/svelte/icons/user";
	import Users from "@lucide/svelte/icons/users";
	import { scale } from "svelte/transition";
	import { quintOut } from "svelte/easing";

	// Initialize from localStorage on mount
	$effect(() => {
		formStore.initFromLocalStorage();
	});
</script>

<div class="mx-auto flex w-full max-w-2xl flex-col items-center px-4 py-8">
	{#if wizardStore.teamType === "individual"}
		<div class="neo-border neo-shadow-lg mb-8 -rotate-1 bg-primary px-6 py-3">
			<div class="flex items-center gap-3">
				<User class="size-6" />
				<h1 class="text-2xl font-black tracking-tight uppercase md:text-3xl">
					Your Information
				</h1>
			</div>
		</div>

		<p class="mb-8 max-w-md text-center text-base font-medium">
			We'll save this locally so you don't have to enter it again.
		</p>

		<div class="w-full space-y-6">
			<div
				class="neo-border neo-shadow space-y-4 bg-card p-6"
				in:scale={{ duration: 300, delay: 100, easing: quintOut }}
			>
				<div class="space-y-2">
					<Label for="student-name" class="text-sm font-bold uppercase">Full Name</Label>
					<Input
						id="student-name"
						name="studentName"
						type="text"
						placeholder="Enter your full name"
						bind:value={formStore.studentName}
						class="neo-border-sm py-3 text-lg"
					/>
				</div>

				<div class="space-y-2">
					<Label for="reg-number" class="text-sm font-bold uppercase"
						>Registration Number</Label
					>
					<Input
						id="reg-number"
						name="regNo"
						type="text"
						placeholder="Enter your registration number"
						bind:value={formStore.regNo}
						class="neo-border-sm py-3 text-lg"
					/>
				</div>
			</div>
		</div>
	{:else}
		<div class="neo-border neo-shadow-lg mb-8 rotate-1 bg-secondary px-6 py-3">
			<div class="flex items-center gap-3">
				<Users class="size-6" />
				<h1 class="text-2xl font-black tracking-tight uppercase md:text-3xl">
					Group Members
				</h1>
			</div>
		</div>

		<p class="mb-8 max-w-md text-center text-base font-medium">
			Add all the members of your group (minimum 2, maximum {STUDENT_LIMIT}).
		</p>

		<div class="w-full space-y-4">
			{#each formStore.students as student, index (student.id)}
				<div
					class="neo-border neo-shadow relative bg-card p-4"
					in:scale={{ duration: 300, delay: index * 50, easing: quintOut }}
					out:scale={{ duration: 200, easing: quintOut }}
				>
					<div class="mb-4 flex items-center justify-between">
						<div class="flex items-center gap-2">
							<div class="neo-border-sm bg-accent px-3 py-1">
								<span class="text-sm font-black">#{index + 1}</span>
							</div>
						</div>
						<Button
							type="button"
							variant="outline"
							size="sm"
							onclick={() => formStore.removeStudent(student.id)}
							disabled={formStore.students.length <= 2}
							class="neo-border-sm hover:bg-destructive hover:text-white"
						>
							<X class="h-4 w-4" />
						</Button>
					</div>
					<div class="grid gap-4 sm:grid-cols-2">
						<div class="space-y-2">
							<Label
								for="student-{student.id}-name"
								class="text-xs font-bold uppercase"
							>
								Full Name
							</Label>
							<Input
								id="student-{student.id}-name"
								type="text"
								placeholder="Enter full name"
								bind:value={student.name}
								oninput={(e) =>
									formStore.updateStudent(
										student.id,
										"name",
										(e.target as HTMLInputElement).value
									)}
								required
								class="neo-border-sm"
							/>
						</div>
						<div class="space-y-2">
							<Label
								for="student-{student.id}-regNo"
								class="text-xs font-bold uppercase"
							>
								Registration Number
							</Label>
							<Input
								id="student-{student.id}-regNo"
								type="text"
								placeholder="Enter reg number"
								bind:value={student.regNo}
								oninput={(e) =>
									formStore.updateStudent(
										student.id,
										"regNo",
										(e.target as HTMLInputElement).value
									)}
								required
								class="neo-border-sm"
							/>
						</div>
					</div>
				</div>
			{/each}

			{#if formStore.students.length < STUDENT_LIMIT}
				<button
					type="button"
					onclick={formStore.addStudent}
					class="neo-border flex w-full cursor-pointer items-center justify-center gap-2 border-dashed bg-muted p-4 text-sm font-bold uppercase transition-colors hover:bg-card"
					in:scale={{
						duration: 300,
						delay: formStore.students.length * 50,
						easing: quintOut
					}}
				>
					<Plus class="h-5 w-5" />
					Add Another Student
				</button>
			{/if}
		</div>
	{/if}
</div>
