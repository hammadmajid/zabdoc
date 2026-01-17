import { formStore } from "./form-store.svelte";

export type DocumentType = "Assignment" | "Lab Task" | null;
export type TeamType = "individual" | "group" | null;

export type WizardStep =
    | "select-document"
    | "select-team"
    | "student-info"
    | "course-info"
    | "document-info"
    | "images"
    | "finalize";

function createWizardStore() {
    let currentStep = $state<WizardStep>("select-document");
    let documentType = $state<DocumentType>(null);
    let teamType = $state<TeamType>(null);
    let direction = $state<"forward" | "backward">("forward");

    // Determine the step order based on document type
    const stepOrder = $derived.by(() => {
        const baseSteps: WizardStep[] = [
            "select-document",
            "select-team",
            "student-info",
            "course-info",
            "document-info",
        ];

        // Add images step only for Lab Task
        if (documentType === "Lab Task") {
            baseSteps.push("images");
        }

        baseSteps.push("finalize");
        return baseSteps;
    });

    const currentStepIndex = $derived(stepOrder.indexOf(currentStep));
    const totalSteps = $derived(stepOrder.length);
    const progress = $derived(((currentStepIndex + 1) / totalSteps) * 100);
    const isFirstStep = $derived(currentStepIndex === 0);
    const isLastStep = $derived(currentStepIndex === totalSteps - 1);

    function setDocumentType(type: DocumentType) {
        documentType = type;
        if (type) {
            // Sync with form store
            formStore.setDocumentType(type);
            // Set default marks based on document type
            if (type === "Assignment") {
                formStore.setDocumentMarks("4");
            } else if (type === "Lab Task") {
                formStore.setDocumentMarks("7.5");
            }
        }
    }

    function setTeamType(type: TeamType) {
        teamType = type;
        if (type) {
            formStore.isMultiMode = type === "group";
        }
    }

    function nextStep() {
        const steps = stepOrder;
        const currentIndex = steps.indexOf(currentStep);
        if (currentIndex < steps.length - 1) {
            direction = "forward";
            currentStep = steps[currentIndex + 1];
        }
    }

    function prevStep() {
        const steps = stepOrder;
        const currentIndex = steps.indexOf(currentStep);
        if (currentIndex > 0) {
            direction = "backward";
            currentStep = steps[currentIndex - 1];
        }
    }

    function goToStep(step: WizardStep) {
        const steps = stepOrder;
        const targetIndex = steps.indexOf(step);
        const currentIndex = steps.indexOf(currentStep);

        if (targetIndex !== -1) {
            direction = targetIndex > currentIndex ? "forward" : "backward";
            currentStep = step;
        }
    }

    function reset() {
        currentStep = "select-document";
        documentType = null;
        teamType = null;
        direction = "forward";
        formStore.reset();
    }

    function canProceed(): boolean {
        switch (currentStep) {
            case "select-document":
                return documentType !== null;
            case "select-team":
                return teamType !== null;
            case "student-info":
                if (teamType === "individual") {
                    return formStore.studentName.trim() !== "" && formStore.regNo.trim() !== "";
                } else {
                    return formStore.students.every(s => s.name.trim() !== "" && s.regNo.trim() !== "");
                }
            case "course-info":
                return formStore.selectedClass !== "" && formStore.selectedCourse !== "";
            case "document-info":
                return formStore.document.number !== "" && formStore.document.date !== "";
            case "images":
                return formStore.hasImages === true;
            case "finalize":
                return true;
            default:
                return false;
        }
    }

    return {
        get currentStep() { return currentStep; },
        get documentType() { return documentType; },
        get teamType() { return teamType; },
        get direction() { return direction; },
        get currentStepIndex() { return currentStepIndex; },
        get totalSteps() { return totalSteps; },
        get progress() { return progress; },
        get isFirstStep() { return isFirstStep; },
        get isLastStep() { return isLastStep; },
        get stepOrder() { return stepOrder; },

        setDocumentType,
        setTeamType,
        nextStep,
        prevStep,
        goToStep,
        reset,
        canProceed,
    };
}

export const wizardStore = createWizardStore();
