import React, {useState, useEffect} from 'react'

export default function Home() {
    const [data, setData] = useState(null);
    const [error, setError] = useState(null);
    useEffect(() => {
        getData()
    }, []);

    const getData = async () => {
        try {
            // const res = await fetch("https://backendtry1.herokuapp.com/")
            const res = await fetch(process.env.NEXT_PUBLIC_BACKEND_URL);
            const data = await res.json();
            setData(data);
        }
        catch (error) {
          setError(error);
        }
    }
    if (error) {
        return <div>Failed to load {error.toString()}</div>
    }
    //tidak ke
    if (!data) {
        return <div>Loading...</div>
    }
    if (!data.data){
        return <div>Data kosong</div>
    }
    return (
        <div>
            {data.data.map((item, index) => (
                <p key={index}>{item}</p>
            ))}
            {/* <p>data: {data.data}</p> */}
        </div>
    )
}
