import { Link } from "react-router-dom";

function Home() {
  return (
    <>
      <section className="center">
        <h1>Bienvenue sur MIRANDE</h1>
        <Link to="/login" type="button" className="counter">
          Se connecter
        </Link>
      </section>
    </>
  );
}

export default Home;
