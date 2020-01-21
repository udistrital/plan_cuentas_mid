package documentopresupuestalmanager

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/udistrital/plan_cuentas_mid/models"
	"github.com/udistrital/utils_oas/formatdata"
	"github.com/udistrital/utils_oas/request"
)

var documentoPrespuestalURI = beego.AppConfig.String("financieraMongoCurdApiService") + "documento_presupuestal" + "/"

func GetAllPresupuestalDocumentFromCRUDByType(vigencia, cg, docType string) ([]models.DocumentoPresupuestal, error) {
	var rows []models.DocumentoPresupuestal
	var err error
	var responseData map[string]interface{}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()
	err = request.GetJson(documentoPrespuestalURI+vigencia+"/"+cg+"/"+docType, &responseData)
	if err == nil {
		err = formatdata.FillStruct(responseData["Body"], &rows)
	}
	return rows, err
}

func GetAllPresupuestalDocumentFromCRUDByMovParentUUID(vigencia, cg, docUUID string) ([]models.DocumentoPresupuestal, error) {
	var rows []models.DocumentoPresupuestal
	var err error
	var responseData map[string]interface{}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()
	err = request.GetJson(documentoPrespuestalURI+"get_doc_mov_by_parent/"+vigencia+"/"+cg+"/"+docUUID, &responseData)
	if err == nil {
		err = formatdata.FillStruct(responseData["Body"], &rows)
	}
	return rows, err
}

func GetOnePresupuestalDocumentFromCRUDByID(vigencia, cg, UUID string) (models.DocumentoPresupuestal, error) {
	var row models.DocumentoPresupuestal
	var err error
	var responseData map[string]interface{}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s", r)
		}
	}()
	err = request.GetJson(documentoPrespuestalURI+"documento/"+vigencia+"/"+cg+"/"+UUID, &responseData)
	if err == nil {
		err = formatdata.FillStruct(responseData["Body"], &row)
	}
	return row, err
}
