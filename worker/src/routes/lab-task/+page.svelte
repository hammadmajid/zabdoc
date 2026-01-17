<script lang="ts">
    import Form from "$lib/components/form.svelte";
    import StudentForm from "$lib/components/forms/student.svelte";
    import CourseForm from "$lib/components/forms/course.svelte";
    import ImagesForm from "$lib/components/forms/images.svelte";
    import DueDate from "$lib/components/forms/fields/due-date.svelte";
    import SEO from "$lib/components/seo.svelte";
    import * as Card from "$lib/components/ui/card/index";
    import { Input } from "$lib/components/ui/input";
    import { Label } from "$lib/components/ui/label";
    import { Separator } from "$lib/components/ui/separator";
    import { formStore } from "$lib/stores/form-store.svelte";
    import ClipboardList from "@lucide/svelte/icons/clipboard-list";

    // Set document type on mount
    $effect(() => {
        formStore.setDocumentType("Lab Task");
        formStore.setDocumentMarks("7.5");
    });
</script>

<SEO
    title="lab task | zabdoc"
    description="Generate lab task PDFs for SZABIST students."
    canonical="https://zabdoc.xyz/lab-task"
    url="https://zabdoc.xyz/lab-task"
/>

<div class="space-y-12">
    <div class="space-y-2">
        <div class="flex items-center gap-3">
            <div class="neo-border-sm neo-shadow-sm bg-accent p-2">
                <ClipboardList class="size-6" />
            </div>
            <h3 class="text-3xl font-black uppercase tracking-tight">
                Lab Task
            </h3>
        </div>
        <p class="font-medium">
            Fill out the information below to generate the lab task.
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
                        >Information about the lab task.</Card.Description
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
                            placeholder="7.5"
                            value={formStore.document.marks}
                            oninput={(e) =>
                                formStore.setDocumentMarks(
                                    (e.target as HTMLInputElement).value,
                                )}
                        />
                    </div>

                    <div class="space-y-2">
                        <Label for="document-number">Number</Label>
                        <Input
                            id="document-number"
                            name="number"
                            type="number"
                            placeholder="1-15"
                            required
                            value={formStore.document.number}
                            oninput={(e) =>
                                formStore.setDocumentNumber(
                                    (e.target as HTMLInputElement).value,
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

        <Separator />
        <ImagesForm />
    </Form>
</div>
