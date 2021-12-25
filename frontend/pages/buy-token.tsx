import type { NextPage } from 'next'
import Head from 'next/head'
import { useState } from 'react'
import styles from '../styles/Home.module.css'
import Web3 from 'web3'
import { AbiItem } from 'web3-utils'
import sellKEEYTokenABI from '../contracts/sellKEEYToken.json'
import tetherABI from '../contracts/tether.json'
import contractAddresses from '../contracts/contractAddress'

declare const window: any

const BuyToken: NextPage = () => {
    const [tokenNumber, setTokenNumber] = useState<number>(0)
    const [remainingTokenNumber, setRemainingTokenNumber] = useState<number>(0)
    const [currentAddress, setCurrentAddress] = useState<string | null>(null)

    const decimals = 18

    const detectProvider: any = () => {
        if (typeof window !== undefined) {
            if (window.ethereum !== undefined) {
                return window.ethereum;
            } else if (typeof window.web3 !== undefined) {
                return window.web3.currentProvder
            }
        }

        throw 'No provider found'
    }

    // @todo: load remaining KEEY tokens number
    // const loadRemainingToken = async (): Promise<number> => {
    //     const keeyToken = await new web3.eth.Contract(
    //         tetherABI as AbiItem[],
    //         '0xdb113b28Fd4E7E772F56cfE938AebA165BD2C47A'
    //     )
    // }

    const initWeb3 = async (): Promise<Web3> => {
        const provider = detectProvider()

        return new Web3(provider)
    }

    const getCurrentMetamaskAcc = async (web3: Web3): Promise<any> => {
        const accounts = await web3.eth.getAccounts()

        if (accounts[0] !== currentAddress) {
            setCurrentAddress(accounts[0])

            return accounts[0]
        }
    }

    const buyKeeyToken = async () => {
        const web3 = await initWeb3()
        const currentAcc = await getCurrentMetamaskAcc(web3)

        const numberWithDecimals = tokenNumber * 10 ** decimals

        const usdtToken = await new web3.eth.Contract(
            tetherABI as AbiItem[],
            contractAddresses.usdt
        )
        const contract = await new web3.eth.Contract(
            sellKEEYTokenABI as AbiItem[],
            contractAddresses.sellKEEYToken
        )

        const amountToPay = await contract.methods
            .getAmountToPay(tokenNumber)
            .call()
        const amountToPayInDecimals = BigInt(amountToPay * 10 ** decimals)
        const allowance = await usdtToken.methods.allowance(
            currentAcc,
            contractAddresses.sellKEEYToken
        )
            .call()

        if (allowance < amountToPayInDecimals) {
            await usdtToken.methods
                .approve(
                    contractAddresses.sellKEEYToken,
                    amountToPayInDecimals.toString()
                )
                .send({ from: currentAcc })
        }

        await contract.methods.buyKEEYToken(numberWithDecimals.toString()).send({
            from: currentAcc
        })
    }

    return (
        <div className={styles.container}>
            <Head>
                <title>Bootcamp</title>
                <meta name="description" content="Bootcamp" />
            </Head>

            <main className={styles.main}>
                <h1>
                    {!currentAddress && "Not connected"}
                    {currentAddress && "Welcome " + currentAddress}
                </h1>
                <h1>Remaining tokens: {remainingTokenNumber}</h1>
                <div className={styles.grid}>
                    <button name='buy-token' onClick={buyKeeyToken} className={styles.card}>
                        Buy token
                    </button>
                </div>
                <input type="text" id="buy-token" onChange={(e) => setTokenNumber(parseFloat(e.target.value))} />
            </main>

            <footer className={styles.footer}>
                KAI Bootcamp
            </footer>
        </div>
    )
}

export default BuyToken
