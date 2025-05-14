import Typograph from "@/components/Typograph";
import TestService from "@/services/test-service";

const TestView = async () => {
  const tests = await TestService.list();

  return (
    <div>
      <Typograph.H1>Tests</Typograph.H1>

      <div className="flex flex-col border border-black rounded-lg gap-2">
        {tests.map((t) => (
          <div key={t.id} className="flex">
            <Typograph.Span>{t.name}</Typograph.Span>
          </div>
        ))}
      </div>
    </div>
  );
};

export default TestView;
