type Props = {
  children: React.ReactNode;
};

const Container = ({ children }: Props) => (
  <div className="w-full h-full p-4 bg-background min-w-screen min-h-screen">
    {children}
  </div>
);

export default Container;
