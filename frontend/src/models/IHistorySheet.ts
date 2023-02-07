import { OutpatientScreeningsInterface } from "./OutpatientScreening";
import { NursesInterface } from "./INurse";

/* เอา Field มาจาก Backend เชค type ให้ตรงด้วย (type วิธีเขียนไม่เหมือนกันนะ)

*/

export interface HistorySheetsInterface {
    ID: number;
    Room: string;

    HistorySheetTimeStart: Date;
    HistorySheetTimeStop: Date;

    MemberID: number;
    Member: NursesInterface;     //มันเป็น Object

    FoodOrdereds: OutpatientScreeningsInterface[];  //Interface FoodOrdered มันรับเป็น Array
}
//historySheet