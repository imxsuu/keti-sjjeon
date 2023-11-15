package main

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
	"time"
)

var Wait_goFunc sync.WaitGroup

func check_ImagePullBackOff() {

	defer Wait_goFunc.Done()

	var podName string

	for {

		cmd_checkImagePullBackOff := exec.Command("bash", "-c", "kubectl get pods -n gedge-platform | grep ImagePullBackOff | awk '{print $1}' | head -1")
		output_checkImagePullBackOff, err := cmd_checkImagePullBackOff.Output()
		if err != nil {
			fmt.Println(err)
		}

		str_checkImagePullBackOff := string(output_checkImagePullBackOff)
		podName = str_checkImagePullBackOff
		podName = strings.ReplaceAll(podName, "\n", "")

		if podName == "" {
			continue
		} else {
			fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] GEdge 응용 " + podName + "이 ImagePullBackOff 상태로 감지되었습니다.")
			time.Sleep(time.Second * 3)
			retCmd := exec.Command("bash", "-c", "kubectl delete pod --grace-period=0 --force --namespace gedge-platform "+podName)
			_, err := retCmd.Output()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "]" + " ImagePullBackOff 상태로 감지된 GEdge 응용 " + podName + "을 제거하였습니다.\n")
		}
	}
}

func check_OOMKilled() {

	defer Wait_goFunc.Done()

	var podName string

	for {
		cmd_checkOOMKilled := exec.Command("bash", "-c", "kubectl get pods -n gedge-platform | grep OOMKilled | awk '{print $1}' | head -1")
		output_checkOOMKilled, err := cmd_checkOOMKilled.Output()
		if err != nil {
			fmt.Println(err)
		}

		str_checkOOMKilled := string(output_checkOOMKilled)
		podName = str_checkOOMKilled
		podName = strings.ReplaceAll(podName, "\n", "")

		if podName == "" {
			continue
		} else {
			fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] GEdge 응용 " + podName + "이 OOMKilled 상태로 감지되었습니다.")
			time.Sleep(time.Second * 3)
			retCmd := exec.Command("bash", "-c", "kubectl delete pod --grace-period=0 --force --namespace gedge-platform "+podName)
			_, err := retCmd.Output()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "]" + " OOMKilled 상태로 감지된 GEdge 응용 " + podName + "을 제거하였습니다.\n")
		}
	}
}

func check_Pending() {

	defer Wait_goFunc.Done()

	var podName string

	for {

		//fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] Check Pending status....")
		time.Sleep(time.Second * 2)

		cmd_checkPending := exec.Command("bash", "-c", "kubectl get pods -n gedge-platform | grep Pending | awk '{print $1}' | head -1")
		output_checkPending, err := cmd_checkPending.Output()
		if err != nil {
			fmt.Println(err)
		}

		str_checPending := string(output_checkPending)
		podName = str_checPending
		podName = strings.ReplaceAll(podName, "\n", "")

		if podName == "" {
			continue
		} else {
			fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] GEdge 응용 " + podName + "이 Pending 상태로 감지되었습니다.")
			time.Sleep(time.Second * 3)
			retCmd := exec.Command("bash", "-c", "kubectl delete pod --grace-period=0 --force --namespace gedge-platform "+podName)
			_, err := retCmd.Output()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "]" + " Pending 상태로 감지된 GEdge 응용 " + podName + "을 제거하였습니다.\n")
		}
	}
}

func check_Error() {

	defer Wait_goFunc.Done()

	var podName string

	for {

		//fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] Check Error status....")
		time.Sleep(time.Second * 2)

		cmd_checkError := exec.Command("bash", "-c", "kubectl get pods -n gedge-platform | grep Error | awk '{print $1}' | head -1")
		output_checkError, err := cmd_checkError.Output()
		if err != nil {
			fmt.Println(err)
		}

		str_checkError := string(output_checkError)
		podName = str_checkError
		podName = strings.ReplaceAll(podName, "\n", "")

		if podName == "" {
			continue
		} else {
			fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] GEdge 응용 " + podName + "이 Error 상태로 감지되었습니다.")
			time.Sleep(time.Second * 3)
			//retCmd := exec.Command("bash", "-c", "kubectl delete pod --grace-period=0 --force --namespace gedge-platform " + podName)
			//_, err := retCmd.Output()
			//if err != nil {
			//        fmt.Println(err)
			//}
			//fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "]" + " Error 상태로 감지된 GEdge 응용 " + podName + "을 제거하였습니다.\n")
		}
	}
}

func check_CrashLoopBackOff() {

	defer Wait_goFunc.Done()

	var podName string

	for {

		cmd_checkCrashLoopBackOff := exec.Command("bash", "-c", "kubectl get pods -n gedge-platform | grep CrashLoopBackOff | awk '{print $1}' | head -1")
		output_checkCrashLoopBackOff, err := cmd_checkCrashLoopBackOff.Output()
		if err != nil {
			fmt.Println(err)
		}

		str_checkCrashLoopBackOff := string(output_checkCrashLoopBackOff)
		podName = str_checkCrashLoopBackOff
		podName = strings.ReplaceAll(podName, "\n", "")

		if podName == "" {
			continue
		} else {
			fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "] GEdge 응용 " + podName + "이 CrashLoopBackOff 상태로 감지되었습니다.")
			time.Sleep(time.Second * 3)
			retCmd := exec.Command("bash", "-c", "kubectl delete pod --grace-period=0 --force --namespace gedge-platform "+podName)
			_, err := retCmd.Output()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("[" + time.Now().Format("2006-01-02 15:04:05") + "]" + " CrashLoopBackOff 상태로 감지된 GEdge 응용 " + podName + "을 제거하였습니다.\n")
		}
	}
}

func main() {

	Wait_goFunc.Add(1)

	go check_ImagePullBackOff()
	go check_OOMKilled()
	//go check_Running()
	//go check_Complete()
	go check_Pending()
	//go check_Terminating()
	//go check_gpuAllocation()
	go check_Error()
	//go check_Complete_logBackup()
	go check_CrashLoopBackOff()

	Wait_goFunc.Wait()
}
