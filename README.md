# RAMid
Projeto de middleware adaptativo da disciplina de Plataformas de Distribuição do CIn-UFPE


# Para criar o plugin
go build -buildmode=plugin -o ../plugins/componentA_v1.so

go build -ldflags "-pluginpath=componentB1" -buildmode=plugin -o ../components/componentB1.so componentB1.go
go build -ldflags "-pluginpath=componentB2" -buildmode=plugin -o ../components/componentB2.so componentB1.go


go build -ldflags "-pluginpath=componentA" -buildmode=plugin -o componentA1.so componentA1.go


go build -ldflags "-pluginpath=plugin/hot-$(date +%s)" -buildmode=plugin -o ../components/componentA1.so componentA1.go
go build -ldflags "-pluginpath=plugin/hot-$(date +%s)" -buildmode=plugin -o ../components/componentB1.so componentB1.go
go build -ldflags "-pluginpath=plugin/hot-$(date +%s)" -buildmode=plugin -o ../components/componentC1.so componentC1.go
go build -ldflags "-pluginpath=plugin/hot-$(date +%s)" -buildmode=plugin -o ../components/componentD1.so componentD1.go


go build -ldflags "-pluginpath=componentA1" -buildmode=plugin -o ../../componentA1.so componentA1.go