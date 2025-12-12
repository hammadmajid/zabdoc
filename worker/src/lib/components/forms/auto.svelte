<script lang="ts">
    import * as Select from "$lib/components/ui/select/index.js";
    import Input from "$lib/components/ui/input/input.svelte";
    import Label from "$lib/components/ui/label/label.svelte";
    import * as Card from "$lib/components/ui/card/index.js";
    import { Switch } from "$lib/components/ui/switch/index.js";
    import Button from "$lib/components/ui/button/button.svelte";
    import { browser } from "$app/environment";
    import { data, type DataStructure } from "$lib/data/data";
    import { toast } from "svelte-sonner";
    import Plus from "@lucide/svelte/icons/plus";
    import X from "@lucide/svelte/icons/x";
    import { scale } from "svelte/transition";
    import { quintOut } from "svelte/easing";

    // LocalStorage utility functions
    function setLocalStorage(key: string, value: string) {
        if (!browser) return;
        try {
            localStorage.setItem(key, value);
        } catch (error) {
            toast.warning("Failed to save to localStorage", {
                description: error as string,
                position: "top-center",
            });
        }
    }

    function getLocalStorage(key: string): string {
        if (!browser) return "";
        try {
            return localStorage.getItem(key) || "";
        } catch (error) {
            toast.warning("Failed to read from localStorage", {
                description: error as string,
                position: "top-center",
            });
            return "";
        }
    }

    interface Student {
        id: number;
        name: string;
        regNo: string;
    }

    let selectedClass = $state("" as keyof DataStructure | "");
    let selectedCourse = $state("");
    let studentName = $state("");
    let regNum = $state("");
    let isMultiMode = $state(false);
    let students = $state<Student[]>([
        { id: 1, name: "", regNo: "" },
        { id: 2, name: "", regNo: "" },
    ]);
    let studentCounter = $state(2);

    const STUDENT_LIMIT = 6;

    // Initialize values from localStorage on component mount
    $effect(() => {
        if (browser) {
            studentName = getLocalStorage("studentName") || "";
            regNum = getLocalStorage("regNum") || "";
            selectedClass =
                (getLocalStorage("class") as keyof DataStructure) || "";
            isMultiMode = getLocalStorage("isMultiMode") === "true";
        }
    });

    // Update localStorage when values change
    $effect(() => {
        if (browser && studentName) {
            setLocalStorage("studentName", studentName);
        }
    });

    $effect(() => {
        if (browser && regNum) {
            setLocalStorage("regNum", regNum);
        }
    });

    $effect(() => {
        if (browser && selectedClass) {
            setLocalStorage("class", selectedClass);
        }
    });

    $effect(() => {
        if (browser) {
            setLocalStorage("isMultiMode", isMultiMode.toString());
        }
    });

    function addStudent() {
        if (students.length < 6) {
            studentCounter++;
            students = [
                ...students,
                { id: studentCounter, name: "", regNo: "" },
            ];
        }
    }

    function removeStudent(studentId: number) {
        if (students.length > 2) {
            students = students.filter((student) => student.id !== studentId);
        }
    }

    function updateStudent(
        studentId: number,
        field: "name" | "regNo",
        value: string,
    ) {
        students = students.map((student) =>
            student.id === studentId ? { ...student, [field]: value } : student,
        );
    }

    const classes = $derived(
        Object.keys(data).map((className) => ({
            value: className,
            label: className,
        })),
    );

    const courses = $derived(
        selectedClass && selectedClass in data
            ? Object.keys(data[selectedClass]).map((courseName) => ({
                  value: courseName,
                  label: courseName,
              }))
            : [],
    );

    const courseDetails = $derived(
        selectedClass &&
            selectedCourse &&
            selectedClass in data &&
            selectedCourse in data[selectedClass]
            ? data[selectedClass][selectedCourse]
            : null,
    );

    const classTriggerContent = $derived(
        classes.find((c) => c.value === selectedClass)?.label ?? "Select class",
    );

    const courseTriggerContent = $derived(
        courses.find((c) => c.value === selectedCourse)?.label ??
            "Select course",
    );

    // Reset course selection when class changes
    $effect(() => {
        if (selectedClass) {
            selectedCourse = "";
        }
    });
</script>

<Card.Root class="">
    <Card.Header>
        <div class="flex items-center justify-between">
            <div>
                <Card.Title>Student</Card.Title>
                <Card.Description>This information will be cached locally.</Card.Description>
            </div>
         <div class="flex items-center gap-2">
                <Label for="multi-toggle" class="text-sm">Group</Label>
                <Switch id="multi-toggle" bind:checked={isMultiMode} />
            </div>
        </div>
    </Card.Header>
    <Card.Content class="space-y-4">
        {#if !isMultiMode}
            <div class="space-y-2">
                <Label for="student-name">Full Name</Label>
                <Input
                    id="student-name"
                    name="studentName"
                    type="text"
                    placeholder="Enter your full name"
                    bind:value={studentName}
                />
            </div>

            <div class="space-y-2">
                <Label for="reg-number">Registration Number</Label>
                <Input
                    id="reg-number"
                    name="regNo"
                    type="number"
                    placeholder="Enter your registration number"
                    bind:value={regNum}
                />
            </div>
        {:else}
            <div class="space-y-3">
                {#each students as student, index (student.id)}
                    <div
                        in:scale={{ duration: 400, easing: quintOut }}
                        out:scale={{ duration: 300, easing: quintOut }}
                    >
                        <Card.Root class="relative">
                            <Card.Header>
                                <div class="flex items-center justify-between">
                                    <Card.Title class="text-sm font-medium"
                                        >Student {index + 1}</Card.Title
                                    >
                                    <Button
                                        type="button"
                                        variant="outline"
                                        size="sm"
                                        onclick={() =>
                                            removeStudent(student.id)}
                                        disabled={students.length <= 2}
                                    >
                                        <X class="w-3 h-3" />
                                    </Button>
                                </div>
                            </Card.Header>
                            <Card.Content class="space-y-3">
                                <div class="space-y-2">
                                    <Label
                                        for="student-{student.id}-name"
                                        class="text-xs font-medium"
                                        >Full Name</Label
                                    >
                                    <Input
                                        id="student-{student.id}-name"
                                        name="student-{student.id}-name"
                                        type="text"
                                        placeholder="Enter full name"
                                        bind:value={student.name}
                                        oninput={(e) =>
                                            updateStudent(
                                                student.id,
                                                "name",
                                                (e.target as HTMLInputElement)
                                                    .value,
                                            )}
                                        required
                                    />
                                </div>
                                <div class="space-y-2">
                                    <Label
                                        for="student-{student.id}-regNo"
                                        class="text-xs font-medium"
                                        >Registration Number</Label
                                    >
                                    <Input
                                        id="student-{student.id}-regNo"
                                        name="student-{student.id}-regNo"
                                        type="text"
                                        placeholder="Enter registration number"
                                        bind:value={student.regNo}
                                        oninput={(e) =>
                                            updateStudent(
                                                student.id,
                                                "regNo",
                                                (e.target as HTMLInputElement)
                                                    .value,
                                            )}
                                        required
                                    />
                                </div>
                            </Card.Content>
                        </Card.Root>
                    </div>
                {/each}
                {#if students.length < STUDENT_LIMIT}
                    <div
                        in:scale={{
                            duration: 400,
                            easing: quintOut,
                            delay: 100,
                        }}
                    >
                        <Card.Root class="border-dashed">
                            <Card.Content
                                class="flex items-center justify-center py-6"
                            >
                                <Button
                                    type="button"
                                    variant="outline"
                                    class="w-full"
                                    onclick={addStudent}
                                >
                                    <Plus class="w-4 h-4 mr-2" />
                                    Add Student
                                </Button>
                            </Card.Content>
                        </Card.Root>
                    </div>
                {/if}
            </div>
        {/if}

        <!-- Hidden inputs for multi-mode -->
        {#if isMultiMode}
            <input type="hidden" name="isMultiMode" value="true" />
            {#each students as student, index}
                <input
                    type="hidden"
                    name="student-{index + 1}-name"
                    value={student.name}
                />
                <input
                    type="hidden"
                    name="student-{index + 1}-regNo"
                    value={student.regNo}
                />
            {/each}
        {:else}
            <input type="hidden" name="isMultiMode" value="false" />
        {/if}

        <div class="space-y-2">
            <Label for="class-select">Class</Label>
            <!-- Class Selection -->
            <Select.Root
                type="single"
                name="class"
                required
                bind:value={selectedClass}
            >
                <Select.Trigger
                    id="class-select"
                    class="!w-full whitespace-normal min-h-[2.25rem] h-auto"
                >
                    <span class="text-left leading-tight"
                        >{classTriggerContent}</span
                    >
                </Select.Trigger>
                <Select.Content>
                    <Select.Group>
                        <Select.Label>Available Classes</Select.Label>
                        {#each classes as classItem (classItem.value)}
                            <Select.Item
                                value={classItem.value}
                                label={classItem.label}
                            >
                                {classItem.label}
                            </Select.Item>
                        {/each}
                    </Select.Group>
                </Select.Content>
            </Select.Root>
        </div>
    </Card.Content>
</Card.Root>

<Card.Root class="">
    <Card.Header>
        <Card.Title>Course</Card.Title>
        <Card.Description>
            Instructor and course code are auto filled based on the course.
        </Card.Description>
    </Card.Header>
    <Card.Content class="space-y-4">
        <div class="space-y-2">
            <Label for="course-select">Course</Label>
            <!-- Course Selection -->
            <Select.Root
                type="single"
                name="course"
                required
                bind:value={selectedCourse}
                disabled={!selectedClass}
            >
                <Select.Trigger
                    id="course-select"
                    class="!w-full whitespace-normal min-h-[2.25rem] h-auto"
                >
                    <span class="text-left leading-tight"
                        >{courseTriggerContent}</span
                    >
                </Select.Trigger>
                <Select.Content>
                    <Select.Group>
                        <Select.Label>Available Courses</Select.Label>
                        {#each courses as course (course.value)}
                            <Select.Item
                                value={course.value}
                                label={course.label}
                            >
                                {course.label}
                            </Select.Item>
                        {/each}
                    </Select.Group>
                </Select.Content>
            </Select.Root>
            {#if !selectedClass}
                <p class="text-xs text-muted-foreground">
                    Please select a class first to see available courses.
                </p>
            {/if}
        </div>

        <!-- Auto-filled fields -->
        {#if courseDetails}
            <div class="space-y-2">
                <Label for="instructor">Instructor</Label>
                <Input
                    id="instructor"
                    name="instructor"
                    type="text"
                    placeholder="Instructor"
                    value={courseDetails.instructor}
                    readonly
                    class="bg-muted"
                />
            </div>

            <div class="space-y-2">
                <Label for="course-code">Course Code</Label>
                <Input
                    id="course-code"
                    name="courseCode"
                    type="text"
                    placeholder="Course Code"
                    value={courseDetails.code}
                    readonly
                    class="bg-muted"
                />
            </div>
        {/if}
    </Card.Content>
</Card.Root>
