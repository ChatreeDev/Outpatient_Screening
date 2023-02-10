import { OutpatientScreeningsInterface } from "./IOutpatientScreening";

export interface DiabetesLevelInterface {
    ID: number;
    Level: string;
    AssessmentForms: string;
    HistoryTakingForms: string;

    OutpatientScreenings: OutpatientScreeningsInterface[];
}