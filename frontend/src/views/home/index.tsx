import { Link } from "@tanstack/react-router";

const Home = () => {
  return (
    <div className="items-center justify-center flex flex-col h-screen">
      <h1>Titulo Titular</h1>

      <Link to="/tests">
        <span>Teste</span>
      </Link>
    </div>
  );
};

export default Home;
