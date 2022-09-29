import React, {useState, useEffect} from 'react'
interface getData{
    data?: Array<Dictionary>
}
interface Dictionary {
    [key: string]: any
} 
export default function Home() {
    const [data, setData] = useState<getData>(null);
    const [load, setLoad] = useState(false)
    const [error, setError] = useState(null);
    const [showModal, setShowModal] = useState(false)
    const [username, setUsername] = useState("")
    const [password, setPassword] = useState("")
    const [token, setToken] = useState(null)
    useEffect(() => {
        getData()
    }, []);
    useEffect(()=>{
        getData()
    }, [token])
    const getData = async () => {
        try {
            console.log(token)
            // const res = await fetch("https://backendtry1.herokuapp.com/")
            const res = await fetch(process.env.NEXT_PUBLIC_BACKEND_URL,{
                headers:{
                    "Authorization":token
                }
            });
            const data = await res.json();
            if (res.status != 200){
                throw Error(data.error)
            }
            console.log(data)
            setData(data);
            setError(null)
        }
        catch (error) {
            setError(error.message);
        }
    }
    const checked = async (i)=>{
        setLoad(true)
        let newArray = [...data.data]

        let newData = newArray[i]
        let updateData = {
            "check":(newData.done?"0":"1"),
            "id":newData.ID.toString()
        }
        try {
            
            // const res = await fetch("https://backendtry1.herokuapp.com/")
            const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/update-check`, {
                method: 'POST',
                body: JSON.stringify(updateData),
                headers:{
                    "Authorization":token
                }
            });
            const data = await res.json();
            newData.done = (!newData?.done) ?? false
            newArray[i] = newData
            setData({...data, data:newArray})
        }
        catch (error) {
            setError(error);
        }
        setLoad(false)
        
    }
    const deleteData = async(i)=>{
        setLoad(true)
        let newArray = [...data.data]
        let newData = newArray[i]
        let deleteData = {
            "id":newData.ID.toString()
        }
        try {
            
            const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/delete-todo`, {
                method: 'POST',
                body: JSON.stringify(deleteData),
                headers:{
                    "Authorization":token
                }
            });
            getData()
        }
        catch (error) {
            setError(error);
        }
        setLoad(false)
        
    }
    const login = async()=>{
        try{
            let res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/login`,{
                method:'POST',
                body:JSON.stringify({'username':username, "password":password})
            })
            let mes= await res.json()
            if(res.status != 200){
                alert("Failed Login")
                return
            }
            console.log(mes.token)
            setToken(mes.token)
            setError(null)
            setShowModal(false)
            // getData()
        }catch(error){

        }
        
    }
    return (
        <div className="pt-3">
            
            {showModal ? (
                <>
                <div
                    className="justify-center items-center flex overflow-x-hidden overflow-y-auto fixed inset-0 z-50 outline-none focus:outline-none"
                >
                    <div className="relative  my-6 mx-auto max-w-3xl w-1/2">
                    {/*content*/}
                    <div className="border-0 rounded-lg shadow-lg relative flex flex-col w-full bg-white outline-none focus:outline-none">
                        {/*header*/}
                        <div className="flex items-start justify-between p-5 border-b border-solid border-slate-200 rounded-t">
                        <h3 className="text-3xl font-semibold">
                            Login
                        </h3>
                        <button
                            className="p-1 ml-auto bg-transparent border-0 text-black opacity-5 float-right text-3xl leading-none font-semibold outline-none focus:outline-none"
                            onClick={() => setShowModal(false)}
                        >
                            <span className="bg-transparent text-black opacity-5 h-6 w-6 text-2xl block outline-none focus:outline-none">
                            ×
                            </span>
                        </button>
                        </div>
                        {/*body*/}
                        <div className="relative p-6 flex-auto">
                            <div className="mb-4">
                                <label className="block text-gray-700 text-sm font-bold mb-2" >
                                    Username
                                </label>
                                <input value={username} onChange={(e)=>setUsername(e.target.value)} className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" id="username" type="text" placeholder="Username"/>
                                </div>
                                <div className="mb-6">
                                <label className="block text-gray-700 text-sm font-bold mb-2" >
                                    Password
                                </label>
                                <input value={password} onChange={(e)=>setPassword(e.target.value)} className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" id="password" type="password" placeholder="******************"/>
                               
                            </div>
                            <button onClick={login}>Login</button>
                        </div>
                        {/*footer*/}
                        <div className="flex items-center justify-end p-6 border-t border-solid border-slate-200 rounded-b">
                        <button
                            className="text-red-500 background-transparent font-bold uppercase px-6 py-2 text-sm outline-none focus:outline-none mr-1 mb-1 ease-linear transition-all duration-150"
                            type="button"
                            onClick={() => setShowModal(false)}
                        >
                            Close
                        </button>
                        
                        </div>
                    </div>
                    </div>
                </div>
                <div className="opacity-25 fixed inset-0 z-40 bg-black"></div>
                </>
            ) : null}

            <h1 className="text-center font-bold text-blue text-3xl">Welcome to Record</h1>
            <div className="p-3 border w-1/2 m-auto" >
            {token == null? <button className="bg-green-400 rounded mb-3 px-2 py-1" onClick={()=>setShowModal(true)}>Login</button> : <button className="bg-green-400 rounded mb-3 px-2 py-1" onClick={()=>{setToken(null);setData(null);setError(null)}}>Logout</button>} 
            <Input onSuccess={getData} token={token}/>
            </div>
            <div className="pl-10">
            {error ?
                <div>Failed to load {error.toString()}</div>
                :
                !data || load ?
                    <div>Loading ....</div>
                    : (
                        (data?.data ?? []).length === 0 ? <p>Data kosong</p>
                        : (data?.data?? [ ]).map((data, i)=>< div key={i}>
                        <p >ID: {data.ID} task: {data.task} 
                        <input type="checkbox" className="mx-2" defaultChecked={data.done} onChange={()=>checked(i)} />
                        <button className="bg-red-700 px-2 text-white rounded" onClick={()=>deleteData(i)}>x</button>
                         </p>
                        
                        </div>)
                    )
            }
            </div>


        </div>
    )
}
function listTask(){
    return(
        <>
            
        </>
    )
}
function Input({onSuccess, token}) {
    const [val, setVal] = useState('')
    const [data, setData] = useState(null);
    const [error, setError] = useState(null);
    const handleSubmit = async (e) => {
        e.preventDefault()
        if(val === ""){
            alert("data harus dimasukkan")
            return 
        }
        const body = {
          task: val
        }
    
        try {
            const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/send`, {
                method: 'POST',
                body: JSON.stringify(body),
                headers:{
                    "Authorization":token
                }
            });
            const data = await res.json();
            setData(data.message);
            setVal('')
            onSuccess();
        }
        catch (error) {
            setError(error);
        }
    }
    return (
        <div>
            <form onSubmit={handleSubmit} >
            <div className="flex">
            <input type="text" id="default-input" placeholder="input data" className="bg-gray-50 border border-gray-300 
                text-gray-900 text-sm rounded-lg focus:ring-blue-500 w-full
            focus:border-blue-500 block p-2.5 dark:bg-gray-700 dark:border-gray-600 
            dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
            onChange={(e)=>setVal(e.target.value)} value={val}
            />
            <button className="bg-blue-400 text-white  rounded-lg ml-1 p-2.5">Submit</button>
            </div>
            
            </form>
        </div>
    )
}

