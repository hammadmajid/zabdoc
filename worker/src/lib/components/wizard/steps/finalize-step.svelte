<script lang="ts">
    import { wizardStore } from "$lib/stores/wizard-store.svelte";
    import { formStore } from "$lib/stores/form-store.svelte";
    import { smartName } from "$lib/utils";
    import Button from "$lib/components/ui/button/button.svelte";
    import Loader2 from "@lucide/svelte/icons/loader-2";
    import CheckCircle from "@lucide/svelte/icons/check-circle";
    import XCircle from "@lucide/svelte/icons/x-circle";
    import FileText from "@lucide/svelte/icons/file-text";
    import User from "@lucide/svelte/icons/user";
    import Users from "@lucide/svelte/icons/users";
    import FileQuestion from "@lucide/svelte/icons/file-question";
    import BookOpen from "@lucide/svelte/icons/book-open";
    import CalendarDays from "@lucide/svelte/icons/calendar-days";
    import ImageIcon from "@lucide/svelte/icons/image";
    import Pencil from "@lucide/svelte/icons/pencil";
    import Download from "@lucide/svelte/icons/download";
    import RefreshCw from "@lucide/svelte/icons/refresh-cw";
    import { scale } from "svelte/transition";
    import { quintOut } from "svelte/easing";

    let isLoading = $state(false);
    let isSuccess = $state(false);
    let isError = $state(false);
    let errorMessage = $state("");

    async function handleSubmit() {
        isLoading = true;
        isSuccess = false;
        isError = false;
        errorMessage = "";

        try {
            const formData = formStore.buildFormData();

            const apiUrl = window.location.host.includes("localhost")
                ? "http://localhost:8080/generate"
                : "https://api.zabdoc.xyz/generate";

            const response = await fetch(apiUrl, {
                method: "POST",
                body: formData,
            });

            if (!response.ok) {
                const errorText = await response.text().catch(() => "");
                throw new Error(errorText || `Server error: ${response.status}`);
            }

            const blob = await response.blob();
            const url = window.URL.createObjectURL(blob);
            const a = document.createElement("a");
            a.style.display = "none";
            a.href = url;
            a.download = smartName(formData);
            document.body.appendChild(a);
            a.click();
            window.URL.revokeObjectURL(url);
            document.body.removeChild(a);

            isSuccess = true;
        } catch (error) {
            isError = true;
            errorMessage = error instanceof Error ? error.message : String(error);
        }
        isLoading = false;
    }

    function handleStartOver() {
        isSuccess = false;
        isError = false;
        errorMessage = "";
        wizardStore.reset();
    }

    function handleRetry() {
        isError = false;
        errorMessage = "";
        handleSubmit();
    }
</script>

<div class="flex flex-col items-center px-4 py-8 max-w-3xl mx-auto w-full">
    <div class="neo-border neo-shadow-lg bg-primary px-6 py-3 mb-8 rotate-[-1deg]">
        <div class="flex items-center gap-3">
            <CheckCircle class="size-6" />
            <h1 class="text-2xl md:text-3xl font-black uppercase tracking-tight">
                Review & Generate
            </h1>
        </div>
    </div>

    <p class="text-base font-medium max-w-md mb-8 text-center">
        Make sure everything looks correct before generating your PDF.
    </p>

    {#if isSuccess}
        <div
            class="w-full neo-border neo-shadow-lg bg-secondary p-8 text-center space-y-6"
            in:scale={{ duration: 400, easing: quintOut }}
        >
            <div class="neo-border-sm bg-card p-4 inline-block mx-auto">
                <CheckCircle class="size-16 text-green-600" />
            </div>
            <h2 class="text-2xl font-black uppercase">Success!</h2>
            <p class="text-muted-foreground">
                Your PDF has been generated and downloaded.
            </p>
            <div class="flex flex-col sm:flex-row gap-4 justify-center pt-4">
                <Button
                    onclick={handleSubmit}
                    class="neo-border neo-shadow bg-primary hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none transition-all"
                >
                    <Download class="size-5 mr-2" />
                    Download Again
                </Button>
                <Button
                    variant="outline"
                    onclick={handleStartOver}
                    class="neo-border neo-shadow hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none transition-all"
                >
                    Start Over
                </Button>
            </div>
        </div>
    {:else if isError}
        <div
            class="w-full neo-border neo-shadow-lg bg-destructive/10 p-8 text-center space-y-6"
            in:scale={{ duration: 400, easing: quintOut }}
        >
            <div class="neo-border-sm bg-card p-4 inline-block mx-auto border-destructive">
                <XCircle class="size-16 text-destructive" />
            </div>
            <h2 class="text-2xl font-black uppercase text-destructive">Error!</h2>
            <div class="space-y-2">
                <p class="text-muted-foreground">
                    Something went wrong while generating your PDF.
                </p>
                <div class="neo-border-sm bg-card p-3 text-left">
                    <p class="text-sm font-mono text-destructive break-all">
                        {errorMessage}
                    </p>
                </div>
            </div>
            <div class="flex flex-col sm:flex-row gap-4 justify-center pt-4">
                <Button
                    onclick={handleRetry}
                    class="neo-border neo-shadow bg-primary hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none transition-all"
                >
                    <RefreshCw class="size-5 mr-2" />
                    Try Again
                </Button>
                <Button
                    variant="outline"
                    onclick={handleStartOver}
                    class="neo-border neo-shadow hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none transition-all"
                >
                    Start Over
                </Button>
            </div>
        </div>
    {:else}
        <div class="w-full space-y-4">
            <!-- Document Type -->
            <div
                class="neo-border neo-shadow bg-card p-4 flex items-center justify-between"
                in:scale={{ duration: 300, delay: 50, easing: quintOut }}
            >
                <div class="flex items-center gap-3">
                    <div class="neo-border-sm bg-primary p-2">
                        <FileText class="size-5" />
                    </div>
                    <div>
                        <p class="text-xs font-bold uppercase text-muted-foreground">Document Type</p>
                        <p class="font-bold">{wizardStore.documentType}</p>
                    </div>
                </div>
                <button
                    type="button"
                    onclick={() => wizardStore.goToStep("select-document")}
                    class="neo-border-sm bg-muted p-2 hover:bg-secondary transition-colors cursor-pointer"
                >
                    <Pencil class="size-4" />
                </button>
            </div>

            <!-- Student Info -->
            <div
                class="neo-border neo-shadow bg-card p-4 flex items-center justify-between"
                in:scale={{ duration: 300, delay: 100, easing: quintOut }}
            >
                <div class="flex items-center gap-3">
                    <div class="neo-border-sm bg-secondary p-2">
                        {#if wizardStore.teamType === "blank"}
                            <FileQuestion class="size-5" />
                        {:else if wizardStore.teamType === "individual"}
                            <User class="size-5" />
                        {:else}
                            <Users class="size-5" />
                        {/if}
                    </div>
                    <div>
                        <p class="text-xs font-bold uppercase text-muted-foreground">
                            {#if wizardStore.teamType === "blank"}
                                Student Info
                            {:else if wizardStore.teamType === "individual"}
                                Student
                            {:else}
                                Group Members
                            {/if}
                        </p>
                        {#if wizardStore.teamType === "blank"}
                            <p class="font-bold">Left blank</p>
                            <p class="text-sm text-muted-foreground">Fill in after printing</p>
                        {:else if wizardStore.teamType === "individual"}
                            <p class="font-bold">{formStore.studentName || "Not set"}</p>
                            <p class="text-sm text-muted-foreground">{formStore.regNo || "No reg number"}</p>
                        {:else}
                            <p class="font-bold">{formStore.students.length} members</p>
                            <p class="text-sm text-muted-foreground">
                                {formStore.students.map(s => s.name || "Unnamed").join(", ")}
                            </p>
                        {/if}
                    </div>
                </div>
                <button
                    type="button"
                    onclick={() => wizardStore.goToStep("student-info")}
                    class="neo-border-sm bg-muted p-2 hover:bg-secondary transition-colors cursor-pointer"
                >
                    <Pencil class="size-4" />
                </button>
            </div>

            <!-- Course Info -->
            <div
                class="neo-border neo-shadow bg-card p-4 flex items-center justify-between"
                in:scale={{ duration: 300, delay: 150, easing: quintOut }}
            >
                <div class="flex items-center gap-3">
                    <div class="neo-border-sm bg-accent p-2">
                        <BookOpen class="size-5" />
                    </div>
                    <div>
                        <p class="text-xs font-bold uppercase text-muted-foreground">Course</p>
                        <p class="font-bold">{formStore.selectedCourse || "Not selected"}</p>
                        <p class="text-sm text-muted-foreground">
                            {formStore.selectedClass || "No class"} • {formStore.courseDetails?.instructor || "No instructor"}
                        </p>
                    </div>
                </div>
                <button
                    type="button"
                    onclick={() => wizardStore.goToStep("course-info")}
                    class="neo-border-sm bg-muted p-2 hover:bg-secondary transition-colors cursor-pointer"
                >
                    <Pencil class="size-4" />
                </button>
            </div>

            <!-- Document Details -->
            <div
                class="neo-border neo-shadow bg-card p-4 flex items-center justify-between"
                in:scale={{ duration: 300, delay: 200, easing: quintOut }}
            >
                <div class="flex items-center gap-3">
                    <div class="neo-border-sm bg-primary p-2">
                        <CalendarDays class="size-5" />
                    </div>
                    <div>
                        <p class="text-xs font-bold uppercase text-muted-foreground">Details</p>
                        <p class="font-bold">
                            {wizardStore.documentType} #{formStore.document.number || "?"} • {formStore.document.marks || "?"} marks
                        </p>
                        <p class="text-sm text-muted-foreground">
                            Due: {formStore.document.date || "Not set"}
                        </p>
                    </div>
                </div>
                <button
                    type="button"
                    onclick={() => wizardStore.goToStep("document-info")}
                    class="neo-border-sm bg-muted p-2 hover:bg-secondary transition-colors cursor-pointer"
                >
                    <Pencil class="size-4" />
                </button>
            </div>

            <!-- Images (Lab Task only) -->
            {#if wizardStore.documentType === "Lab Task"}
                <div
                    class="neo-border neo-shadow bg-card p-4 flex items-center justify-between"
                    in:scale={{ duration: 300, delay: 250, easing: quintOut }}
                >
                    <div class="flex items-center gap-3">
                        <div class="neo-border-sm bg-secondary p-2">
                            <ImageIcon class="size-5" />
                        </div>
                        <div>
                            <p class="text-xs font-bold uppercase text-muted-foreground">Images</p>
                            <p class="font-bold">
                                {formStore.imageItems.length} image{formStore.imageItems.length !== 1 ? "s" : ""} attached
                            </p>
                        </div>
                    </div>
                    <button
                        type="button"
                        onclick={() => wizardStore.goToStep("images")}
                        class="neo-border-sm bg-muted p-2 hover:bg-secondary transition-colors cursor-pointer"
                    >
                        <Pencil class="size-4" />
                    </button>
                </div>
            {/if}
        </div>

        <!-- Generate Button -->
        <div class="w-full mt-8" in:scale={{ duration: 300, delay: 300, easing: quintOut }}>
            <button
                type="button"
                onclick={handleSubmit}
                disabled={isLoading}
                class="w-full neo-border neo-shadow-lg bg-primary text-primary-foreground px-8 py-6 text-xl font-black uppercase tracking-wide flex items-center justify-center gap-3 hover:translate-x-[4px] hover:translate-y-[4px] hover:shadow-none transition-all cursor-pointer disabled:opacity-50 disabled:cursor-not-allowed disabled:hover:translate-x-0 disabled:hover:translate-y-0 disabled:hover:shadow-[var(--neo-shadow-lg)]"
            >
                {#if isLoading}
                    <Loader2 class="size-6 animate-spin" />
                    Generating...
                {:else}
                    <Download class="size-6" />
                    Generate PDF
                {/if}
            </button>
        </div>
    {/if}
</div>
