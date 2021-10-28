package appdynamics

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"strconv"
	"strings"
	"testing"
	//"github.com/k0kubun/pp"
)

func TestAccAppDHealthRule_basicSingleMetricAllBtsMultipleCrit(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)

	resourceName := "appdynamics_health_rule.test_all_bts_multiple_criteria"

	entityType := "BUSINESS_TRANSACTION_PERFORMANCE"
	businessTransactionScope := "ALL_BUSINESS_TRANSACTIONS"

	criticalConditionAggregationType := "ANY"
	criticalCriteria  := []map[string]interface{} {
		{
			"name": acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum),
			"shortname": strings.ToUpper(acctest.RandStringFromCharSet(2, acctest.CharSetAlpha)),
			"evaluate_to_true_on_no_data": false,
			"eval_detail_type": "SINGLE_METRIC",
			"metric_aggregation_function": "VALUE",
			"metric_path": "95th Percentile Response Time (ms)",
			"metric_eval_detail_type": "SPECIFIC_TYPE",
			"compare_condition": "GREATER_THAN_SPECIFIC_VALUE",
			"compare_value": 2.4,
			"baseline_name": "baseline_name",
			"baseline_condition": "WITHIN_BASELINE",
			"baseline_unit": "PERCENTAGE",
		},
		{
			"name": acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum),
			"shortname": strings.ToUpper(acctest.RandStringFromCharSet(2, acctest.CharSetAlpha)),
			"evaluate_to_true_on_no_data": false,
			"eval_detail_type": "SINGLE_METRIC",
			"metric_aggregation_function": "VALUE",
			"metric_path": "95th Percentile Response Time (ms)",
			"metric_eval_detail_type": "SPECIFIC_TYPE",
			"compare_condition": "GREATER_THAN_SPECIFIC_VALUE",
			"compare_value": 7.5,
			"baseline_name": "baseline_name",
			"baseline_condition": "WITHIN_BASELINE",
			"baseline_unit": "PERCENTAGE",
		},
	}

	warningConditionAggregationType := "ALL"
	warningCriteria  := []map[string]interface{} {
		{
			"name": acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum),
			"shortname": strings.ToUpper(acctest.RandStringFromCharSet(2, acctest.CharSetAlpha)),
			"evaluate_to_true_on_no_data": false,
			"eval_detail_type": "SINGLE_METRIC",
			"metric_aggregation_function": "VALUE",
			"metric_path": "95th Percentile Response Time (ms)",
			"metric_eval_detail_type": "SPECIFIC_TYPE",
			"compare_condition": "GREATER_THAN_SPECIFIC_VALUE",
			"compare_value": 2.8,
			"baseline_name": "baseline_name",
			"baseline_condition": "WITHIN_BASELINE",
			"baseline_unit": "PERCENTAGE",
		},
	}

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: allBTsHealthRule(strings.Split(resourceName, ".")[1], name, entityType, businessTransactionScope, criticalConditionAggregationType, criticalCriteria, warningConditionAggregationType, warningCriteria),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttr(resourceName, "affected_entity_type", entityType),
					resource.TestCheckResourceAttr(resourceName, "business_transaction_scope", businessTransactionScope),
					resource.TestCheckResourceAttr(resourceName, "critical_criteria.0.baseline_condition", "WITHIN_BASELINE"),
					resource.TestCheckResourceAttr(resourceName, "critical_criteria.1.compare_value", "7.5"),
					resource.TestCheckResourceAttr(resourceName, "warning_criteria.0.baseline_condition", "WITHIN_BASELINE"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckHealthRuleDoesNotExist(resourceName)),
	})
}

func TestAccAppDHealthRule_basicSingleMetricAllBtsSingleCrit(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)

	resourceName := "appdynamics_health_rule.test_all_bts_single_criteria"

	entityType := "BUSINESS_TRANSACTION_PERFORMANCE"
	businessTransactionScope := "ALL_BUSINESS_TRANSACTIONS"

	criticalConditionAggregationType := "ANY"
	criticalCriteria := []map[string]interface{} {
		{
			"name": acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum),
			"shortname": strings.ToUpper(acctest.RandStringFromCharSet(2, acctest.CharSetAlpha)),
			"evaluate_to_true_on_no_data": false,
			"eval_detail_type": "SINGLE_METRIC",
			"metric_aggregation_function": "VALUE",
			"metric_path": "95th Percentile Response Time (ms)",
			"metric_eval_detail_type": "SPECIFIC_TYPE",
			"compare_condition": "GREATER_THAN_SPECIFIC_VALUE",
			"compare_value": 2.4,
			"baseline_name": "baseline_name",
			"baseline_condition": "WITHIN_BASELINE",
			"baseline_unit": "PERCENTAGE",
		},
	}

	criticalCriteria2 := []map[string]interface{} {
		{
			"name": acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum),
			"shortname": strings.ToUpper(acctest.RandStringFromCharSet(2, acctest.CharSetAlpha)),
			"evaluate_to_true_on_no_data": false,
			"eval_detail_type": "SINGLE_METRIC",
			"metric_aggregation_function": "VALUE",
			"metric_path": "95th Percentile Response Time (ms)",
			"metric_eval_detail_type": "SPECIFIC_TYPE",
			"compare_condition": "GREATER_THAN_SPECIFIC_VALUE",
			"compare_value": 7.7,
			"baseline_name": "baseline_name",
			"baseline_condition": "WITHIN_BASELINE",
			"baseline_unit": "PERCENTAGE",
		},
	}

	warningConditionAggregationType := "ALL"
	var warningCriteria []map[string]interface{}
	warningCriteria = nil

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: allBTsHealthRule(strings.Split(resourceName, ".")[1], name, entityType, businessTransactionScope, criticalConditionAggregationType, criticalCriteria, warningConditionAggregationType, warningCriteria),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttr(resourceName, "affected_entity_type", entityType),
					resource.TestCheckResourceAttr(resourceName, "business_transaction_scope", businessTransactionScope),
					resource.TestCheckResourceAttr(resourceName, "critical_criteria.0.metric_eval_detail_type", "SPECIFIC_TYPE"),
					resource.TestCheckResourceAttr(resourceName, "critical_criteria.0.compare_value", "2.4"),
					//resource.TestCheckResourceAttrSet(resourceName, "warning_criteria.0.baseline_condition"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
			{
				Config: allBTsHealthRule(strings.Split(resourceName, ".")[1], name, entityType, businessTransactionScope, criticalConditionAggregationType, criticalCriteria2, warningConditionAggregationType, warningCriteria),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttr(resourceName, "affected_entity_type", entityType),
					resource.TestCheckResourceAttr(resourceName, "business_transaction_scope", businessTransactionScope),
					resource.TestCheckResourceAttr(resourceName, "critical_criteria.0.metric_eval_detail_type", "SPECIFIC_TYPE"),
					resource.TestCheckResourceAttr(resourceName, "critical_criteria.0.compare_value", "2.4"),
					//resource.TestCheckResourceAttrSet(resourceName, "warning_criteria.0.baseline_condition"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckHealthRuleDoesNotExist(resourceName)),
	})
}

/*
func TestAccAppDHealthRule_updateSingleMetricAllBts(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)

	resourceName := "appdynamics_health_rule.test_all_bts"
	aggregationFunction := "VALUE"
	detailType := "SINGLE_METRIC"
	entityType := "BUSINESS_TRANSACTION_PERFORMANCE"
	metric := "95th Percentile Response Time (ms)"
	condition := "GREATER_THAN_SPECIFIC_VALUE"
	warnValue := "1"
	criticalValue := "2"

	updatedAggregationFunction := "SUM"
	updatedMetric := "95th Percentile Response Time (ms)"
	updatedCondition := "LESS_THAN_SPECIFIC_VALUE"
	updatedWarnValue := "3"
	updatedCriticalValue := "4"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: allBTsHealthRule(name, aggregationFunction, detailType, entityType, metric, condition, warnValue, criticalValue),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttr(resourceName, "eval_detail_type", detailType),
					resource.TestCheckResourceAttr(resourceName, "affected_entity_type", entityType),
					resource.TestCheckResourceAttr(resourceName, "business_transaction_scope", "ALL_BUSINESS_TRANSACTIONS"),
					resource.TestCheckResourceAttr(resourceName, "metric_eval_detail_type", "SPECIFIC_TYPE"), //bug in api?
					resource.TestCheckResourceAttr(resourceName, "metric_path", metric),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
			{
				Config: allBTsHealthRule(name, updatedAggregationFunction, detailType, entityType, updatedMetric, updatedCondition, updatedWarnValue, updatedCriticalValue),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttr(resourceName, "eval_detail_type", detailType),
					resource.TestCheckResourceAttr(resourceName, "affected_entity_type", entityType),
					resource.TestCheckResourceAttr(resourceName, "business_transaction_scope", "ALL_BUSINESS_TRANSACTIONS"),
					resource.TestCheckResourceAttr(resourceName, "metric_eval_detail_type", "SPECIFIC_TYPE"), //bug in api?
					resource.TestCheckResourceAttr(resourceName, "metric_path", updatedMetric),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckHealthRuleDoesNotExist(resourceName)),
	})
}

func TestAccAppDHealthRule_basicSpecificBts(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)
	bts := []string{bt1}

	resourceName := "appdynamics_health_rule.test_specific_bts"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: specificBTsHealthRule(name, bts),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "business_transactions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckHealthRuleDoesNotExist(resourceName)),
	})
}

func TestAccAppDHealthRule_updateSpecificBts(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)
	bts := []string{bt1}
	updatedBts := []string{bt1, bt2}

	resourceName := "appdynamics_health_rule.test_specific_bts"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: specificBTsHealthRule(name, bts),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "business_transactions.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
			{
				Config: specificBTsHealthRule(name, updatedBts),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "business_transactions.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckHealthRuleDoesNotExist(resourceName)),
	})
}

func TestAccAppDHealthRule_basicSpecificTiers(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)
	tiers := []string{tier1}

	resourceName := "appdynamics_health_rule.test_specific_tiers"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: specificTiersHealthRule(name, tiers),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "specific_tiers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckHealthRuleDoesNotExist(resourceName)),
	})
}

func TestAccAppDHealthRule_updateSpecificTiers(t *testing.T) {

	name := acctest.RandStringFromCharSet(11, acctest.CharSetAlphaNum)
	tiers := []string{tier1}
	updatedBts := []string{tier1, tier2}

	resourceName := "appdynamics_health_rule.test_specific_tiers"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"appdynamics": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: specificTiersHealthRule(name, tiers),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "specific_tiers.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
			{
				Config: specificTiersHealthRule(name, updatedBts),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "specific_tiers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "application_id", applicationIdS),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					RetryCheck(CheckHealthRuleExists(resourceName)),
				),
			},
		},
		CheckDestroy: RetryCheck(CheckHealthRuleDoesNotExist(resourceName)),
	})
}
*/
func CheckHealthRuleExists(resourceName string) func(state *terraform.State) error {
	return func(state *terraform.State) error {

		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		id, err := strconv.Atoi(resourceState.Primary.ID)
		if err != nil {
			return err
		}

		_, err = appDClient.GetHealthRule(id, applicationIdI)
		if err != nil {
			return err
		}

		return nil
	}
}

func CheckHealthRuleDoesNotExist(resourceName string) func(state *terraform.State) error {
	return func(state *terraform.State) error {

		resourceState, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		id, err := strconv.Atoi(resourceState.Primary.ID)
		if err != nil {
			return err
		}

		_, err = appDClient.GetHealthRule(id, applicationIdI)
		if err == nil {
			return fmt.Errorf("health rule found: %d", id)
		}

		return nil
	}
}

func prepareHealthRuleCondition(level string,
								name string,
								shortname string,
								evaluate_to_true_on_no_data bool,
								eval_detail_type string,
								metric_aggregation_function string,
								metric_path string,
								metric_eval_detail_type string,
								baseline_condition string,
								baseline_name string,
								baseline_unit string,
								compare_value float64,
								compare_condition string) string {
	return fmt.Sprintf(`
            %s_criteria {
                name = "%s"
                shortname = "%s"
                evaluate_to_true_on_no_data = %s
                eval_detail_type = "%s"
                metric_aggregation_function = "%s"
                metric_path = "%s"
                metric_eval_detail_type = "%s"
                baseline_condition = "%s"
                baseline_name = "%s"
                baseline_unit = "%s"
                compare_value = %f
                compare_condition = "%s"
            }`,
	level,
	name,
	shortname,
	strconv.FormatBool(evaluate_to_true_on_no_data),
	eval_detail_type,
	metric_aggregation_function,
	metric_path,
	metric_eval_detail_type,
	baseline_condition,
	baseline_name,
	baseline_unit,
	float64(compare_value),
	compare_condition)
}

func allBTsHealthRule(resourceName string, name string, entityType string, businessTransactionScope string, criticalConditionAggregationType string, criticalCriteria []map[string]interface{}, warningConditionAggregationType string, warningCriteria []map[string]interface{}) string {

	var criticalCriteriaList []string
	var warningCriteriaList []string
	var criticalCriteriaData []string
	var warningCriteriaData []string
	var dataTmp string

	for _, crit := range criticalCriteria {
		criticalCriteriaList = append(criticalCriteriaList, "data.appdynamics_health_rule_critical_condition." + crit["name"].(string))

		dataTmp = prepareHealthRuleCondition( "critical",
												crit["name"].(string),
												crit["shortname"].(string),
												crit["evaluate_to_true_on_no_data"].(bool),
												crit["eval_detail_type"].(string),
												crit["metric_aggregation_function"].(string),
												crit["metric_path"].(string),
												crit["metric_eval_detail_type"].(string),
												crit["baseline_condition"].(string),
												crit["baseline_name"].(string),
												crit["baseline_unit"].(string),
												crit["compare_value"].(float64),
												crit["compare_condition"].(string))

		criticalCriteriaData = append(criticalCriteriaData, dataTmp)
	}

	for _, crit := range warningCriteria {
		warningCriteriaList = append(warningCriteriaList, "data.appdynamics_health_rule_warning_condition." + crit["name"].(string))

		dataTmp = prepareHealthRuleCondition( "warning",
			crit["name"].(string),
			crit["shortname"].(string),
			crit["evaluate_to_true_on_no_data"].(bool),
			crit["eval_detail_type"].(string),
			crit["metric_aggregation_function"].(string),
			crit["metric_path"].(string),
			crit["metric_eval_detail_type"].(string),
			crit["baseline_condition"].(string),
			crit["baseline_name"].(string),
			crit["baseline_unit"].(string),
			crit["compare_value"].(float64),
			crit["compare_condition"].(string))

		warningCriteriaData = append(warningCriteriaData, dataTmp)
	}

	criteriaData := fmt.Sprintf(`
%s

resource "appdynamics_health_rule" "%s" {
  	name = "%s"
  	application_id = var.application_id
  	affected_entity_type = "%s"
	business_transaction_scope = "%s"
	critical_condition_aggregation_type = "%s"
	warning_condition_aggregation_type = "%s"
  	%s
	%s
}`, configureConfig(), resourceName, name, entityType, businessTransactionScope, criticalConditionAggregationType, warningConditionAggregationType, strings.Join(criticalCriteriaData,"\n"), strings.Join(warningCriteriaData, "\n"))

	//fmt.Print(criteriaData)

	return criteriaData
}

//func allBTsHealthRule(name string, aggregationFunction string, detailType string, entityType string, metric string, compareCondition string, warnValue string, criticalValue string) string {
//	return fmt.Sprintf(`
//					%s
//					resource "appdynamics_health_rule" "test_all_bts" {
//					  name = "%s"
//					  application_id = var.application_id
//					  metric_aggregation_function = "%s"
//					  eval_detail_type = "%s"
//					  affected_entity_type = "%s"
//					  business_transaction_scope = "ALL_BUSINESS_TRANSACTIONS"
//					  metric_eval_detail_type = "SPECIFIC_TYPE"
//					  metric_path = "%s"
//					  compare_condition="%s"
//					  warn_compare_value = %s
//					  critical_compare_value = %s
//					}
//`, configureConfig(), name, aggregationFunction, detailType, entityType, metric, compareCondition, warnValue, criticalValue)
//}

/*
func specificBTsHealthRule(name string, bts []string) string {
	return fmt.Sprintf(`
					%s
					resource "appdynamics_health_rule" "test_specific_bts" {
					  name = "%s"
					  application_id = var.application_id
					  metric_aggregation_function = "VALUE"
					  eval_detail_type = "SINGLE_METRIC"
					  affected_entity_type = "BUSINESS_TRANSACTION_PERFORMANCE"
					  business_transaction_scope = "SPECIFIC_BUSINESS_TRANSACTIONS"
					  business_transactions = "%s"
					  metric_eval_detail_type = "SPECIFIC_TYPE"
					  metric_path = "95th Percentile Response Time (ms)"
					  compare_condition = "GREATER_THAN_SPECIFIC_VALUE"
					  warn_compare_value = 100
					  critical_compare_value = 200
					}
`, configureConfig(), name, arrayToString(bts))
}

func specificTiersHealthRule(name string, tiers []string) string {
	return fmt.Sprintf(`
					%s
					resource "appdynamics_health_rule" "test_specific_tiers" {
					  name = "%s"
					  application_id = var.application_id
					  metric_aggregation_function = "VALUE"
					  eval_detail_type = "SINGLE_METRIC"
					  affected_entity_type = "BUSINESS_TRANSACTION_PERFORMANCE"
					  business_transaction_scope = "BUSINESS_TRANSACTIONS_IN_SPECIFIC_TIERS"
					  specific_tiers = %s
					  metric_eval_detail_type = "SPECIFIC_TYPE"
					  metric_path = "95th Percentile Response Time (ms)"
					  compare_condition = "GREATER_THAN_SPECIFIC_VALUE"
					  warn_compare_value = 100
					  critical_compare_value = 200
					}
`, configureConfig(), name, arrayToString(tiers))
}
*/