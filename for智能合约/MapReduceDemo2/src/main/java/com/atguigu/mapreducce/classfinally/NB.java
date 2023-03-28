package com.atguigu.mapreducce.classfinally;

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.File;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

import org.apache.hadoop.fs.Path;

public class NB {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		predictAll();
	}


	private static String modelFilePath = "D://bayes/2020131128_moxing.txt";
	private static String testDataFilePath = "D://bayes/test.txt";
	public static HashMap<String, Integer> parameters = null;
	public static double Nd = 0.;
	public static HashMap<String, Integer> allFeatures = null;
	public static HashMap<String, Double> labelFeatures = null;
	public static HashSet<String> V = null;


	public static void loadModel(String modelFile) throws Exception {
		if (parameters != null && allFeatures != null) {
			return;
		}
		parameters = new HashMap<String, Integer>();
		allFeatures = new HashMap<String, Integer>();
		labelFeatures = new HashMap<String, Double>();
		V = new HashSet<String>();
		BufferedReader br = new BufferedReader(new FileReader(modelFile));
		String line = null;
		while ((line = br.readLine()) != null) {
			String feature = line.substring(0, line.indexOf("\t"));
			Integer count = Integer.parseInt(line.substring(line.indexOf("\t") + 1));
			if (feature.contains("_")) {
				allFeatures.put(feature, count);
				String label = feature.substring(0, feature.indexOf("_"));
				if (labelFeatures.containsKey(label)) {
					labelFeatures.put(label, labelFeatures.get(label) + count);
				} else {
					labelFeatures.put(label, (double) count);
				}
				String word = feature.substring(feature.indexOf("_") + 1);
				if (!V.contains(word)) {
					V.add(word);
				}
			} else {
				parameters.put(feature, count);
				Nd += count;
			}
		}
		br.close();
	}

	public static String predict(String sentence, String modelFile) throws Exception {
		loadModel(modelFile);

		String predLabel = null;
		double maxValue = Double.NEGATIVE_INFINITY;
		String[] words = sentence.split(" ");
		Set<String> labelSet = parameters.keySet();
		for (String label : labelSet) {
			double tempValue = Math.log(parameters.get(label) / Nd);

			for (String word : words) {
				String lf = label + "_" + word;
				// 计算最大似然概率
				if (allFeatures.containsKey(lf)) {
					tempValue += Math.log((double) (allFeatures.get(lf) + 1) / (labelFeatures.get(label) + V.size()));

				} else {
					tempValue += Math.log((double) (1 / (labelFeatures.get(label) + V.size())));
				}
			}
			if (tempValue > maxValue) {
				maxValue = tempValue;
				predLabel = label;
			}
		}
		return predLabel;
	}

	public static void predictAll() {
		double accuracy = 0.;
		int amount = 0;
		try {
			File file = new File("D://bayes/2020131128_yucewenjian.txt");
			file.createNewFile();
			BufferedWriter out = new BufferedWriter(new FileWriter(file));
			List<String> testData = Files.readAllLines(Paths.get(testDataFilePath));
			for (String instance : testData) {
				String gold;
				if(instance.indexOf(":") == -1) {
					gold = instance.substring(0, instance.indexOf("："));
				} else {
					gold = instance.substring(0, instance.indexOf(":"));
				}
				String sentence = instance.substring(instance.indexOf("\t") + 1);
				String prediction = predict(sentence, modelFilePath);
				System.out.println("Gold='" + gold + "'\tPrediction='" + prediction + "'");
				if (gold.equals(prediction)) {
					accuracy += 1;
				}
				amount += 1;
				out.write(amount + "\t" + prediction+"\n");
			}
			out.flush();
			out.close();
		} catch (Exception e) {
			e.printStackTrace();
		}
		System.out.println("Accuracy = " + accuracy / amount);
	}
}