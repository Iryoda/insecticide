export enum ActionType {
  CLICK = "click",
  DOUBLECLICK = "doubleclick",
  INPUT_TEXT = "text",
  GOTO = "goto",
}

export enum AssertionType {
  TEXT = "text",
}

export type Step = {
  id: string;
  name: string;
  isUnique: boolean;
  selector: string;
  componentId: string;
  nextStepId: string;
  action: ActionType;
  assertion: AssertionType;
  assertionTarget: string;
  updatedAt: Date;
  createdAt: Date;
};
