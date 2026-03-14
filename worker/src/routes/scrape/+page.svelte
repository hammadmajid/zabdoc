<script lang="ts">
    import { fade, scale } from "svelte/transition";
    import { quintOut } from "svelte/easing";
    import ArrowLeft from "@lucide/svelte/icons/arrow-left";
    import Database from "@lucide/svelte/icons/database";
    import Loader2 from "@lucide/svelte/icons/loader-2";
    import XCircle from "@lucide/svelte/icons/x-circle";
    import * as Card from "$lib/components/ui/card";
    import * as Tabs from "$lib/components/ui/tabs";
    import * as Table from "$lib/components/ui/table";
    import Button from "$lib/components/ui/button/button.svelte";
    import Input from "$lib/components/ui/input/input.svelte";
    import Label from "$lib/components/ui/label/label.svelte";

    interface AttendanceRecord {
        lecture: string;
        date: string;
        status: string;
    }

    interface CourseAttendance {
        courseName: string;
        instructor: string;
        records: AttendanceRecord[];
    }

    interface MarkEntry {
        head: string;
        max: string;
        obtained: string;
    }

    type MarkCategory = "quiz" | "assignment" | "lab" | "project" | "exam" | "other";

    interface MarkGroup {
        category: MarkCategory;
        label: string;
        entries: MarkEntry[];
        totalMax: number;
        totalObtained: number | null;
    }

    interface CourseData {
        courseName: string;
        instructor: string;
        records: AttendanceRecord[];
        marks: MarkEntry[];
    }

    import loadingMessages from "$lib/loading-msgs";

    let username = $state("");
    let password = $state("");
    let isLoading = $state(false);
    let isError = $state(false);
    let errorMessage = $state("");
    let loadingMessage = $state("");
    let loadingInterval: ReturnType<typeof setInterval> | null = null;
    let courseData = $state<CourseData[]>([]);
    let hasResults = $state(false);

    function getRandomLoadingMessage() {
        return loadingMessages[Math.floor(Math.random() * loadingMessages.length)];
    }

    function startLoadingMessages() {
        loadingMessage = getRandomLoadingMessage();
        loadingInterval = setInterval(() => {
            loadingMessage = getRandomLoadingMessage();
        }, 2000);
    }

    function stopLoadingMessages() {
        if (loadingInterval) {
            clearInterval(loadingInterval);
            loadingInterval = null;
        }
    }

    async function handleScrap() {
        if (!username.trim() || !password.trim()) {
            isError = true;
            errorMessage = "Username and password are required";
            return;
        }

        isLoading = true;
        isError = false;
        errorMessage = "";
        startLoadingMessages();

        try {
            const apiUrl = window.location.host.includes("localhost")
                ? "http://localhost:8080/scrap"
                : "https://api.zabdoc.xyz/scrap";

            const response = await fetch(apiUrl, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({ username, password }),
            });

            if (!response.ok) {
                const errorData: any = await response.json().catch(() => null);
                if (response.status === 400) {
                    throw new Error(errorData?.error || "Invalid username or password");
                } else if (response.status === 500) {
                    throw new Error(errorData?.error || "Server error. Please try again later.");
                } else {
                    throw new Error(errorData?.error || `Error: ${response.status}`);
                }
            }

            const result: any = await response.json();

            console.log("API Response:", result);

            if (!result.success) {
                throw new Error(result.error || "Failed to fetch attendance and marks data");
            }

            // Ensure data is an object: { course_name: { attendence: {}, marks: {} } }
            if (result.data && typeof result.data === "object" && !Array.isArray(result.data)) {
                courseData = Object.entries(result.data as Record<string, any>).map(([courseName, courseValue]) => {
                    const attendence = courseValue?.attendence ?? {};
                    const marks = courseValue?.marks ?? {};

                    return {
                        courseName,
                        instructor: attendence?.instructor || "Unknown Instructor",
                        records: Array.isArray(attendence?.records) ? attendence.records : [],
                        marks: Array.isArray(marks?.entries) ? marks.entries : []
                    };
                });
            } else {
                courseData = [];
            }

            console.log("Processed courseData:", courseData);
            hasResults = true;
        } catch (error) {
            isError = true;
            errorMessage = error instanceof Error ? error.message : String(error);
        } finally {
            stopLoadingMessages();
            loadingMessage = "";
            isLoading = false;
        }
    }

    function getStatusColor(status: string): string {
        const normalizedStatus = status.toLowerCase().trim();
        if (normalizedStatus === "present" || normalizedStatus === "p") {
            return "text-green-600 bg-green-100 dark:bg-green-950";
        } else if (normalizedStatus === "absent" || normalizedStatus === "a") {
            return "text-red-600 bg-red-100 dark:bg-red-950";
        } else if (normalizedStatus === "late" || normalizedStatus === "l") {
            return "text-yellow-600 bg-yellow-100 dark:bg-yellow-950";
        }
        return "text-muted-foreground bg-muted";
    }

    function getMarkCellClass(obtained: string): string {
        if (obtained.toLowerCase() === "not entered") {
            return "text-muted-foreground italic";
        }
        return "";
    }

    function parseMarkValue(value: string): number | null {
        const normalized = value.trim().toLowerCase();
        if (!normalized || normalized === "not entered") {
            return null;
        }

        const parsed = Number.parseFloat(value);
        return Number.isFinite(parsed) ? parsed : null;
    }

    function formatMarkValue(value: number): string {
        if (Number.isInteger(value)) {
            return `${value}`;
        }

        return value.toFixed(2).replace(/\.00$/, "").replace(/(\.\d)0$/, "$1");
    }

    function getMarkCategory(head: string): MarkCategory {
        const normalized = head.trim().toLowerCase();
        if (normalized.includes("quiz")) return "quiz";
        if (normalized.includes("assignment")) return "assignment";
        if (normalized.includes("lab") || normalized.includes("practical")) return "lab";
        if (normalized.includes("project")) return "project";
        if (
            normalized.includes("mid") ||
            normalized.includes("final") ||
            normalized.includes("paper") ||
            normalized.includes("exam")
        ) {
            return "exam";
        }
        return "other";
    }

    function groupMarks(marks: MarkEntry[]): { groups: MarkGroup[]; overallMax: number; overallObtained: number | null } {
        const categoryOrder: MarkCategory[] = ["quiz", "assignment", "lab", "project", "exam", "other"];
        const categoryLabels: Record<MarkCategory, string> = {
            quiz: "Quizzes",
            assignment: "Assignments",
            lab: "Lab Work",
            project: "Projects",
            exam: "Exams",
            other: "Other",
        };

        const groupMap = new Map<MarkCategory, { entries: MarkEntry[]; totalMax: number; totalObtained: number; hasObtained: boolean }>();

        let overallMax = 0;
        let overallObtained = 0;
        let hasOverallObtained = false;

        for (const mark of marks) {
            const category = getMarkCategory(mark.head);
            if (!groupMap.has(category)) {
                groupMap.set(category, { entries: [], totalMax: 0, totalObtained: 0, hasObtained: false });
            }
            const group = groupMap.get(category)!;

            const max = parseMarkValue(mark.max) ?? 0;
            const obtained = parseMarkValue(mark.obtained);

            group.entries.push(mark);
            group.totalMax += max;
            if (obtained !== null) {
                group.totalObtained += obtained;
                group.hasObtained = true;
            }

            overallMax += max;
            if (obtained !== null) {
                overallObtained += obtained;
                hasOverallObtained = true;
            }
        }

        const groups: MarkGroup[] = [];
        for (const category of categoryOrder) {
            const data = groupMap.get(category);
            if (!data) continue;
            groups.push({
                category,
                label: categoryLabels[category],
                entries: data.entries,
                totalMax: data.totalMax,
                totalObtained: data.hasObtained ? data.totalObtained : null,
            });
        }

        return {
            groups,
            overallMax,
            overallObtained: hasOverallObtained ? overallObtained : null,
        };
    }
</script>

<div class="min-h-[70vh] flex flex-col">
    <!-- Back to Home Button -->
    <div class="mb-4" in:fade={{ duration: 300 }}>
        <a
            href="/"
            class="neo-border neo-shadow bg-card px-4 py-2 font-bold uppercase text-sm items-center gap-2 hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none transition-all inline-flex w-fit"
        >
            <ArrowLeft class="size-4" />
            Back to Home
        </a>
    </div>

    {#if !hasResults}
        <!-- Login Form -->
        <div class="flex flex-col items-center justify-center px-4 py-8 flex-1 max-w-md mx-auto w-full">
            <div class="neo-border neo-shadow-lg bg-primary px-8 py-4 mb-8 rotate-[-2deg]">
                <div class="flex items-center gap-3">
                    <Database class="size-6" />
                    <h1 class="text-2xl md:text-3xl font-black uppercase tracking-tight">
                        Scrape Data
                    </h1>
                </div>
            </div>

            <p class="text-base font-medium max-w-md mb-8 text-center">
                Enter your ZabDesk credentials to fetch attendance and marks data.
            </p>

            {#if isError}
                <div
                    class="w-full neo-border neo-shadow-lg bg-destructive/10 p-4 mb-6 border-destructive"
                    in:scale={{ duration: 400, easing: quintOut }}
                >
                    <div class="flex items-center gap-3">
                        <XCircle class="size-5 text-destructive flex-shrink-0" />
                        <p class="text-sm font-medium text-destructive">
                            {errorMessage}
                        </p>
                    </div>
                </div>
            {/if}

            <div class="w-full space-y-4">
                <div class="space-y-2">
                    <Label for="username" class="font-bold uppercase text-sm">Username</Label>
                    <Input
                        id="username"
                        type="text"
                        placeholder="Enter your username"
                        bind:value={username}
                        disabled={isLoading}
                        class="neo-border-sm"
                    />
                </div>

                <div class="space-y-2">
                    <Label for="password" class="font-bold uppercase text-sm">Password</Label>
                    <Input
                        id="password"
                        type="password"
                        placeholder="Enter your password"
                        bind:value={password}
                        disabled={isLoading}
                        class="neo-border-sm"
                        onkeydown={(e) => {
                            if (e.key === "Enter") {
                                handleScrap();
                            }
                        }}
                    />
                </div>

                <Button
                    onclick={handleScrap}
                    disabled={isLoading}
                    class="w-full neo-border neo-shadow bg-primary text-primary-foreground px-6 py-6 text-lg font-black uppercase hover:translate-x-[2px] hover:translate-y-[2px] hover:shadow-none transition-all mt-6"
                >
                        {#if isLoading}
                        <Loader2 class="size-5 mr-2 animate-spin" />
                        <span class="text-sm normal-case font-medium tracking-normal">
                            {loadingMessage || "Scraping Data..."}
                        </span>
                    {:else}
                        <Database class="size-5 mr-2" />
                        Scrape Data
                    {/if}
                </Button>
            </div>
        </div>
    {:else}
        <!-- Results Display -->
        <div class="flex flex-col px-4 py-8 max-w-6xl mx-auto w-full">
            <div class="flex items-center justify-between mb-8">
                <div class="neo-border neo-shadow-lg bg-primary px-6 py-3 rotate-[-1deg]">
                    <h1 class="text-2xl md:text-3xl font-black uppercase tracking-tight">
                        Scrapped Data
                    </h1>
                </div>
            </div>

            {#if courseData.length === 0}
                <div class="text-center py-12">
                    <p class="text-lg font-medium text-muted-foreground">
                        No attendance data found.
                    </p>
                </div>
            {:else}
                <div class="grid gap-6 md:grid-cols-2">
                    {#each courseData as course, index}
                        <div
                            in:scale={{ duration: 300, delay: index * 50, easing: quintOut }}
                        >
                            <Card.Root>
                                <Card.Header>
                                    <Card.Title class="text-xl font-black uppercase">
                                        {course.courseName}
                                    </Card.Title>
                                    <Card.Description class="font-medium">
                                        {course.instructor}
                                    </Card.Description>
                                </Card.Header>
                                <Card.Content>
                                    <Tabs.Root value="attendance" class="w-full">
                                        <Tabs.List class="grid w-full grid-cols-2 neo-border-sm">
                                            <Tabs.Trigger value="attendance" class="font-bold uppercase text-sm">
                                                Attendance
                                            </Tabs.Trigger>
                                            <Tabs.Trigger value="marks" class="font-bold uppercase text-sm">
                                                Marks
                                            </Tabs.Trigger>
                                        </Tabs.List>
                                        <Tabs.Content value="attendance" class="mt-4">
                                            {#if !course.records || course.records.length === 0}
                                                <p class="text-sm text-muted-foreground text-center py-4">
                                                    No attendance records available.
                                                </p>
                                            {:else}
                                                <Table.Root>
                                                    <Table.Header>
                                                        <Table.Row>
                                                            <Table.Head>Lecture</Table.Head>
                                                            <Table.Head>Date</Table.Head>
                                                            <Table.Head>Status</Table.Head>
                                                        </Table.Row>
                                                    </Table.Header>
                                                    <Table.Body>
                                                        {#each course.records as record}
                                                            <Table.Row>
                                                                <Table.Cell class="font-medium">
                                                                    {record.lecture}
                                                                </Table.Cell>
                                                                <Table.Cell class="text-muted-foreground">
                                                                    {record.date}
                                                                </Table.Cell>
                                                                <Table.Cell>
                                                                    <span
                                                                        class="inline-block px-2 py-1 rounded text-xs font-bold uppercase {getStatusColor(record.status)}"
                                                                    >
                                                                        {record.status}
                                                                    </span>
                                                                </Table.Cell>
                                                            </Table.Row>
                                                        {/each}
                                                    </Table.Body>
                                                </Table.Root>
                                            {/if}
                                        </Tabs.Content>
                                        <Tabs.Content value="marks" class="mt-4">
                                            {#if !course.marks || course.marks.length === 0}
                                                <p class="text-sm text-muted-foreground text-center py-4">
                                                    No marks records available.
                                                </p>
                                            {:else}
                                                {@const { groups, overallMax, overallObtained } = groupMarks(course.marks)}
                                                <Table.Root>
                                                    <Table.Header>
                                                        <Table.Row>
                                                            <Table.Head>Marks Head</Table.Head>
                                                            <Table.Head>Max</Table.Head>
                                                            <Table.Head>Obtained</Table.Head>
                                                        </Table.Row>
                                                    </Table.Header>
                                                    <Table.Body>
                                                        {#each groups as group}
                                                            <!-- Category section header -->
                                                            <Table.Row class="bg-primary/10 hover:bg-primary/10">
                                                                <Table.Cell colspan={3} class="font-black uppercase text-xs tracking-widest py-2 text-primary">
                                                                    {group.label}
                                                                </Table.Cell>
                                                            </Table.Row>
                                                            <!-- Individual entries -->
                                                            {#each group.entries as mark}
                                                                <Table.Row>
                                                                    <Table.Cell class="font-medium pl-6">
                                                                        {mark.head}
                                                                    </Table.Cell>
                                                                    <Table.Cell class="text-muted-foreground">
                                                                        {mark.max}
                                                                    </Table.Cell>
                                                                    <Table.Cell class={getMarkCellClass(mark.obtained)}>
                                                                        {mark.obtained}
                                                                    </Table.Cell>
                                                                </Table.Row>
                                                            {/each}
                                                            <!-- Category total -->
                                                            <Table.Row class="bg-muted/60 border-t">
                                                                <Table.Cell class="font-bold text-xs uppercase tracking-wide pl-6 text-muted-foreground">
                                                                    {group.label} Total
                                                                </Table.Cell>
                                                                <Table.Cell class="font-bold">
                                                                    {formatMarkValue(group.totalMax)}
                                                                </Table.Cell>
                                                                <Table.Cell class={group.totalObtained === null ? "text-muted-foreground italic font-bold" : "font-bold"}>
                                                                    {group.totalObtained === null ? "Not Entered" : formatMarkValue(group.totalObtained)}
                                                                </Table.Cell>
                                                            </Table.Row>
                                                        {/each}
                                                        <!-- Overall total -->
                                                        <Table.Row class="bg-primary/15 border-t-2 hover:bg-primary/15">
                                                            <Table.Cell class="font-black uppercase text-xs tracking-widest text-primary">
                                                                Overall Total
                                                            </Table.Cell>
                                                            <Table.Cell class="font-black text-primary">
                                                                {formatMarkValue(overallMax)}
                                                            </Table.Cell>
                                                            <Table.Cell class={overallObtained === null ? "text-muted-foreground italic font-black" : "font-black text-primary"}>
                                                                {overallObtained === null ? "Not Entered" : formatMarkValue(overallObtained)}
                                                            </Table.Cell>
                                                        </Table.Row>
                                                    </Table.Body>
                                                </Table.Root>
                                            {/if}
                                        </Tabs.Content>
                                    </Tabs.Root>
                                </Card.Content>
                            </Card.Root>
                        </div>
                    {/each}
                </div>
            {/if}
        </div>
    {/if}
</div>
