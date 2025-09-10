<script lang="ts">
    import AutoForm from "$lib/components/auto-form.svelte";
    import DueDate from "$lib/components/due-date.svelte";
    import Button from "$lib/components/ui/button/button.svelte";
    import Input from "$lib/components/ui/input/input.svelte";
    import * as Card from "$lib/components/ui/card/index";

    async function handleSubmission() {
        const res = await fetch("http://localhost:8080/assignment", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(
                {
                    "studentName": "John Doe",
                    "regNo": "2025-001",
                    "class": "CS101",
                    "course": "Introduction to Computer Science",
                    "courseCode": "CSC101",
                    "instructor": "Dr. Smith",
                    "number": "1",
                    "date": "2025-09-10"
                })
        });
        const blob = await res.blob();
        const url = URL.createObjectURL(blob);
        window.open(url);
    }
</script>

<form
        onsubmit={handleSubmission}
        class="grid md:grid-cols-2 gap-8"
>
    <AutoForm/>

    <Card.Root class="">
        <Card.Header>
            <Card.Title>Document</Card.Title>
            <Card.Description
            >Information about the assignment.
            </Card.Description
            >
        </Card.Header>
        <Card.Content class="space-y-4">
            <DueDate/>

            <Input type="number" placeholder="Number"/>
        </Card.Content>
    </Card.Root>

    <Card.Root class="">
        <Card.Header>
            <Card.Description
            >By clicking the button below you accept terms and privacy/
            </Card.Description
            >
        </Card.Header>
        <Card.CardContent>
            <Button type="submit" class="w-full">Generate cover</Button>
        </Card.CardContent>
    </Card.Root>
</form>
