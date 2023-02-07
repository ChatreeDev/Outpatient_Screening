//import { DiabetesLevelsInterface } from "./OutpatientScreening";

export interface DiabetesLevelsInterface {
    ID: number;

    Name: string;
    Detail: string;
    Price: number;

    DiabetesLevels: DiabetesLevelsInterface[];
}