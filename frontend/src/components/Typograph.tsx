type Props = {
  className?: string;
  children: React.ReactNode;
};

const H1 = ({ className, children }: Props) => (
  <h1 className={`font-bold text-3xl text-foreground ${className}`}>
    {children}
  </h1>
);

const H2 = ({ className, children }: Props) => (
  <h1 className={`font-semibold text-2xl text-foreground ${className}`}>
    {children}
  </h1>
);

const H3 = ({ className, children }: Props) => (
  <h1 className={`font-semibold text-xl text-foreground ${className}`}>
    {children}
  </h1>
);

const Span = ({ className, children }: Props) => (
  <span className={`font-normal text-base text-foreground ${className}`}>
    {children}
  </span>
);

const Typograph = { H1, H2, H3, Span };

export default Typograph;
