<script lang="ts">
	import AutoForm from "$lib/components/auto-form.svelte";
	import DueDate from "$lib/components/due-date.svelte";
	import Button from "$lib/components/ui/button/button.svelte";
	import Input from "$lib/components/ui/input/input.svelte";
	import * as Card from "$lib/components/ui/card/index";
</script>

<svelte:head>
	<title>assignment | zabdoc</title>
</svelte:head>

<div class="space-y-12 px-4 md:p-0">
	<div class="prose">
		<h1>Assignment</h1>
	</div>
	<form
		action="/api/assignment"
		method="POST"
		enctype="multipart/form-data"
		class="grid md:grid-cols-2 gap-8"
	>
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

		<Card.Root class="">
			<Card.Header>
				<Card.Description
					>By clicking the button below you accept terms and privacy/
				</Card.Description>
			</Card.Header>
			<Card.CardContent>
				<Button type="submit" class="w-full">Generate cover</Button>
			</Card.CardContent>
		</Card.Root>
	</form>
</div>
