package utils

import (
	"archive/zip"
	"bufio"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type fileUtil struct {

}
func (fileUtil fileUtil) DeleteDir(removeDir string)  {
	removeDir=strings.ReplaceAll(removeDir,"\\","/")
	err:=os.RemoveAll(removeDir)
	if err!=nil{
		log.Printf("remove dir %s read ex,err:%s",removeDir,err.Error())
		return
	}
	dirs,err:=os.ReadDir(removeDir)
	if err!=nil{
		log.Printf("remove dir %s read ex,err:%s",removeDir,err.Error())
		return
	}
	for _, dir := range dirs {
		if dir.IsDir(){
			fileUtil.DeleteDir(removeDir+"/"+dir.Name())
		}else{
			os.Remove(removeDir+"/"+dir.Name())
		}

	}
}
func (fileUtil fileUtil) CopyDir(currentDir string,copyDir string)  {
	currentDir=strings.ReplaceAll(currentDir,"\\","/")
	copyDir=strings.ReplaceAll(copyDir,"\\","/")
	dirs,err:=os.ReadDir(currentDir)
	if err!=nil{
		log.Printf("copy dir  dir %s read ex,err:%s",currentDir,err.Error())
		return
	}
	for _, dir := range dirs {
		if dir.IsDir(){
			fileUtil.CopyDir(currentDir+"/"+dir.Name(),copyDir)
		}else{
			os.Rename(currentDir+"/"+dir.Name(),copyDir+"/"+dir.Name())
		}

	}
}


func (fileUtil fileUtil) GetCurrentDir() string {
	dir,err:=os.Getwd()
	if err!=nil{
		log.Printf("GetCurrentDir,err:%s",err.Error())
	}
	return dir
}
//获取当前目录下的文件或目录信息(不包含多级子目录)
func (fileUtil fileUtil) GetCurrentDirFileInfosAndDirInfos() []string {
	dir:=fileUtil.GetCurrentDir()
	if dir==""{
		return nil
	}
	files,err:=ioutil.ReadDir(dir)
	if err!=nil{
		log.Println("GetCurrentDirFileInfosAndDirInfos ,err:%s",err.Error())
		return nil
	}
	strs:=make([]string,len(files))
	for i, file := range files {
		log.Printf("%d:%s",i,file.Name())
		strs[i]=file.Name()
	}
	return strs
}
//获取当前目录下的文件或目录名(不包含多级子目录)
func (fileUtil fileUtil) GetCurrentDirFilesAndDirs() []string {
	dir:=fileUtil.GetCurrentDir()
	if dir==""{
		return nil
	}
	files,err:=filepath.Glob(filepath.Join(dir,"*"))
	if err!=nil{
		log.Println("GetCurrentDirFilesAndDirs ,err:%s",err.Error())
		return nil
	}
	return files
}

//获取当前文件或目录下的所有文件或目录信息(包括子目录)
func (fileUtil fileUtil) GetCurrentDirAllFilesAndAllDirs() []string {
	dir:=fileUtil.GetCurrentDir()
	if dir==""{
		return nil
	}
	strs:=make([]string,10)
	i:=0
	err:=filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		log.Printf("%d:%s",i,info.Name())
		strs[i]=info.Name()
		i=i+1
		return nil
	})
	if err!=nil{
		log.Println("GetCurrentDirAllFilesAndAllDirs ,err:%s",err.Error())
		return nil
	}
	return strs
}
func (fileUtil) GetFileExtension(file string) string {
	if file == "" {
		return ""
	}
	var strs = strings.Split(file, ".")
	if len(strs) > 1 {
		return strs[len(strs)-1]
	}
	return ""
}
func  (fileUtil) CheckFileIsExists(file string)bool  {
	if fi,err:=os.Stat(file);err==nil{
		fi.Size()
		return true
	}
	return  false
}

func (fileUtil)  CheckFileIsDir(file string)bool  {
	fi,err:=os.Stat(file)
	if err!=nil{
		return false
	}
	return  fi.IsDir()
}

func  (fileUtil fileUtil)  AppendFile(file string,buffer []byte)bool  {
	var f *os.File
	var err error
	if fileUtil.CheckFileIsExists(file){
		f, err = os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666) //打开文件

	}else{
		f, err = os.Create(file) //创建文件
	}
	if err!=nil{
		return  false
	}
	defer  f.Close()
	r, err:=f.Write(buffer)
	if err!=nil{
		return false
	}
	f.Sync()
	log.Println("AppendFile %s file  %d byte success ",file,r)
	return  true
}

//error batch write fail
func (fileUtil fileUtil)   HandlerAppendFile(f *os.File,file string,buffer []byte)bool  {
	var err error
	if fileUtil.CheckFileIsExists(file){
		f, err = os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666) //打开文件

	}else{
		f, err = os.Create(file) //创建文件
	}
	if err!=nil{
		return  false
	}
	r, err:=f.Write(buffer)
	if err!=nil{
		return false
	}
	log.Println("HandlerAppendFile %s file  %d byte success ",file,r)
	return  true
}



func (fileUtil)  Append(file string,buffer []byte)bool  {
	err:=ioutil.WriteFile(file,buffer,0666)
	if err!=nil{
		return  false
	}
	log.Println("Append %s file  success ",file)
	return  true
}

func (fileUtil)   WriteFile(file string,buffer []byte)bool  {
	f, err := os.Create(file) //创建文件
	if err!=nil{
		return  false
	}
	defer  f.Close()
	r, err:=f.Write(buffer)
	if err!=nil{
		return false
	}
	f.Sync()
	log.Println("WriteFile %s file  %d byte success ",file,r)
	return  true
}

func (fileUtil)   Write(file string,buffer []byte)bool  {
	f, err := os.Create(file) //创建文件
	if err!=nil{
		return  false
	}
	defer  f.Close()
	w := bufio.NewWriter(f) //创建文件
	r, err:=w.Write(buffer)
	w.Flush()
	if err!=nil{
		return false
	}
	log.Println("wite %s file  %d byte success ",file,r)
	return  true
}

func (fileUtil)  Zip(srcFile string,descZip string) bool{
	zipFile,err:=os.Create(descZip)
	if err!=nil{
		return  false
	}
	defer zipFile.Close()
	archive:=zip.NewWriter(zipFile)
	defer archive.Close()
	err=filepath.Walk(srcFile, func(path string, info fs.FileInfo, err error) error {
		if err!=nil{
			return err
		}
		header,err:=zip.FileInfoHeader(info)
		if err!=nil{
			return err
		}
		header.Name=strings.TrimPrefix(path,filepath.Dir(srcFile)+"/")
		if info.IsDir(){
			header.Name+="/"
		}else{
			header.Method=zip.Deflate
		}
		writer,err:=archive.CreateHeader(header)
		if err!=nil{
			return err
		}
		if !info.IsDir(){
			file,err:=os.Open(path)
			if err!=nil{
				return err
			}
			defer file.Close()
			_,err =io.Copy(writer,file)
		}
		return  err
	})
	if err!=nil{
		return  false
	}
	return  true
}

func (fileUtil)  UnZip(zipFile string,destDir string)bool{
	zipReader,err:=zip.OpenReader(zipFile)
	if err!=nil{
		return  false
	}
	defer  zipReader.Close()
	for _,f:=range zipReader.File{
		fpath:=filepath.Join(destDir,f.Name)
		if f.FileInfo().IsDir(){
			os.MkdirAll(fpath,os.ModePerm)
		}else{
			if err=os.MkdirAll(filepath.Dir(fpath),os.ModePerm);err!=nil{
				return  false
			}
			inFile,err:=f.Open()
			if err!=nil{
				return  false
			}
			defer  inFile.Close()
			outFile,err:=os.OpenFile(fpath,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,f.Mode())
			if err!=nil{
				return  false
			}
			defer  outFile.Close()
			_,err =io.Copy(outFile,inFile)
		}
	}
	return true
}