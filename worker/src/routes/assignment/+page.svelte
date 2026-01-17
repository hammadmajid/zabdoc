<script lang="ts">
    import Form from "$lib/components/form.svelte";
    import StudentForm from "$lib/components/forms/student.svelte";
    import CourseForm from "$lib/components/forms/course.svelte";
    import DueDate from "$lib/components/forms/fields/due-date.svelte";
    import SEO from "$lib/components/seo.svelte";
    import * as Card from "$lib/components/ui/card/index";
    import Input from "$lib/components/ui/input/input.svelte";
    import Label from "$lib/components/ui/label/label.svelte";
    import { Separator } from "$lib/components/ui/separator";
    import { formStore } from "$lib/stores/form-store.svelte";
    import FileText from "@lucide/svelte/icons/file-text";

    // Set document type on mount
    $effect(() => {
        formStore.setDocumentType("Assignment");
        formStore.setDocumentMarks("4");
    });
</script>

<SEO
    title="assignment | zabdoc"
    description="Generate assignment cover sheets for SZABIST students."
    canonical="https://zabdoc.xyz/assignment"
    url="https://zabdoc.xyz/assignment"
/>

<div class="space-y-12">
    <div class="space-y-2">
        <div class="flex items-center gap-3">
            <div class="neo-border-sm neo-shadow-sm bg-secondary p-2">
                <FileText class="size-6" />
            </div>
            <h3 class="text-3xl font-black uppercase tracking-tight">
                Assignment
            </h3>
        </div>
        <p class="font-medium">
            Fill out the information below to generate the assignment cover.
        </p>
    </div>

    <Form>
        <div class="grid md:grid-cols-2 gap-4">
                <StudentForm />
                <CourseForm />

            <Card.Root>
                <Card.Header>
                    <Card.Title>Document</Card.Title>
                    <Card.Description
                        >Information about the assignment.</Card.Description
                    >
                </Card.Header>
                <Card.Content class="space-y-4">
                    <div class="space-y-2">
                        <Label for="marks">Total Marks</Label>
                        <Input
                            id="marks"
                            name="marks"
                            type="number"
                            step="0.5"
                            placeholder="4"
                            value={formStore.document.marks}
                            oninput={(e) =>
                                formStore.setDocumentMarks(
                                    (e.target as HTMLInputElement).value
                                )}
                        />
                    </div>

                    <div class="space-y-2">
                        <Label for="document-number">Number</Label>
                        <Input
                            id="document-number"
                            name="number"
                            type="number"
                            placeholder="1-4"
                            required
                            value={formStore.document.number}
                            oninput={(e) =>
                                formStore.setDocumentNumber(
                                    (e.target as HTMLInputElement).value
                                )}
                        />
                    </div>

                    <div class="space-y-2">
                        <Label for="due-date">Due Date</Label>
                        <DueDate />
                    </div>
                </Card.Content>
            </Card.Root>
        </div>
    </Form>
</div>
