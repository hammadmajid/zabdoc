<script lang="ts">
	import AutoForm from "$lib/components/auto-form.svelte";
	import DueDate from "$lib/components/due-date.svelte";
	import Button from "$lib/components/ui/button/button.svelte";
	import Input from "$lib/components/ui/input/input.svelte";
	import * as Card from "$lib/components/ui/card/index";
	import Loader2 from "@lucide/svelte/icons/loader-2";

	let isLoading = false;

	function handleSubmit() {
		// The form will continue its normal submission behavior
		// The loading state will be reset when the page potentially reloads or the PDF downloads
		isLoading = true;
	}
</script>

<svelte:head>
	<title>assignment | zabdoc</title>
</svelte:head>

<div class="space-y-12 px-4 md:p-0">
	<div>
		<h1
			class="scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight transition-colors first:mt-0"
		>
			Assignment
		</h1>
	</div>
	<form
		action="https://api.zabdoc.xyz/assignment"
		method="POST"
		enctype="multipart/form-data"
		class="space-y-8"
		on:submit={handleSubmit}
	>
		<div class="grid md:grid-cols-2 gap-8">
			<AutoForm />

			<Card.Root class="">
				<Card.Header>
					<Card.Title>Document</Card.Title>
					<Card.Description
						>Information about the assignment.
					</Card.Description>
				</Card.Header>
				<Card.Content class="space-y-4">
					<DueDate />

					<Input
						name="number"
						type="number"
						placeholder="Number"
						required
					/>
				</Card.Content>
			</Card.Root>
		</div>

		<Card.Root class="">
			<Card.Header>
				<Card.Description
					>By clicking the button below you accept the <a
						href="/about">Terms and Privacy Policy</a
					>.
				</Card.Description>
			</Card.Header>
			<Card.CardContent>
				<Button type="submit" class="w-full" disabled={isLoading}>
					{#if isLoading}
						<Loader2 class="mr-2 h-4 w-4 animate-spin" />
						Generating PDF...
					{:else}
						Generate PDF
					{/if}
				</Button>
			</Card.CardContent>
		</Card.Root>
	</form>
</div>
