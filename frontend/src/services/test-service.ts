import { Test } from "@/entities/app-test";

const createTest = (data?: Partial<Test>): Test => ({
  id: (Math.random() * 100).toString(),
  headComponentId: (Math.random() * 100).toString(),
  name: "any",
  createdAt: new Date(),
  ...data,
});

export default class TestService {
  static async getById(id: string) {
    return { id } as Test;
  }

  static async list() {
    return [createTest(), createTest({ name: "any-1" })];
  }
}
